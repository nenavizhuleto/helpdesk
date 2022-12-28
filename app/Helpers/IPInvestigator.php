<?php

namespace App\Helpers;

class IPInvestigator
{
    private $db;
    private $client_address, $client_port;

    private $server_address;


    public $network, $routers, $credentials, $local_address, $MAC;


    public function __construct($db, $client_address, $client_port)
    {
        $this->db = $db;
        $this->client_address = $client_address;
        $this->client_port = $client_port;

        $this->server_address = "{$_SERVER['SERVER_NAME']}:{$_SERVER['SERVER_PORT']}";

        $this->network = $this->findNetwork();
        $this->routers = $this->findRouters();
        $this->credentials = $this->findRouterCredentials();
        $this->local_address = $this->findLocalAddress();
        $this->MAC = $this->findMACAddress();

    }

    private function ipCIDRCheck($IP, $CIDR)
    {
        list($net, $mask) = explode('/', $CIDR);

        $ip_net = ip2long($net);
        $ip_mask = ~((1 << (32 - $mask)) - 1);

        $ip_ip = ip2long($IP);

        return (($ip_ip& $ip_mask) == ($ip_net& $ip_mask));
    }

    private function findNetwork()
    {
        $network = null;
        $netmask = 32;
        while (!$network && $netmask >= 0) {

            $client_networks = $this->db->query("
            SELECT * 
            FROM `client_networks` 
            WHERE `network` LIKE('%/{$netmask}')
            ")->getResultArray();



            $netmask -= 8;

            if (empty($client_networks)) {
                continue;
            }

            foreach ($client_networks as $client_network) {
                if ($this->ipCIDRCheck($this->client_address, $client_network['network'])) {
                    $network = $client_network;
                    break;
                }
            }
        }

        return $network;
    }

    private function findRouters()
    {
        if (empty($this->network))
            return null;

        $routers = $this->db->query("
        SELECT * 
        FROM client_routers
        WHERE network_id = {$this->network['id']}
        ")->getResultArray();
        return $routers;


    }

    private function findRouterCredentials()
    {
        if (empty($this->routers))
            return null;

        $router_credentials = null;
        foreach ($this->routers as $router) {
            $credentials = $this->db->query("
            SELECT *
            FROM router_credentials
            WHERE router_id = {$router['id']}
            ")->getRowArray();

            if ($credentials['gateway']) {
                $router_credentials = $credentials;
            }

            if ($credentials['address'] == $this->client_address) {
                $router_credentials = $credentials;
                break;
            }
        }


        return $router_credentials;
    }

    private function execSSHCommand($cmd)
    {
        if (empty($this->credentials))
            return null;

        $ssh_cmd = "ssh -l {$this->credentials['login']} -p {$this->credentials['port']} {$this->credentials['address']} '{$cmd}'";
        exec($ssh_cmd, $output);

        return $output;
    }

    private function findLocalAddress()
    {


        $local_address = null;
        $retry_count = 0;

        while ($local_address == null && $retry_count <= 3) {
            $retry_count++;
            $cmd = "ip firewall connection print where dst-address=\"{$this->server_address}\"";
            $output = $this->execSSHCommand($cmd);

            if (empty($output)) {
                continue;
            }



            foreach ($output as $line) {
                if (str_contains($line, "{$this->server_address}") && str_contains($line, "{$this->client_port}")) {
                    $success = preg_match($ip_address_regex, $line, $match);
                    if ($success) {
                        $local_address = explode(':', $match[0])[0];
                    }
                }
            }
        }



        return $local_address;

    }

    private function findMACAddress()
    {
        $MAC = null;

        $cmd = "ip arp print value-list where address=\"{$this->local_address}\"";
        $output = $this->execSSHCommand($cmd);

        if (empty($output)) {
            return $MAC;
        }

        foreach ($output as $line) {
            if (str_contains($line, "mac-address")) {
                $MAC = trim(explode(":", $line, 2)[1]);
            }
        }

        return $MAC;
    }



}