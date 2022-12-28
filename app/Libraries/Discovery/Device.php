<?php

namespace App\Libraries\Discovery;

use ErrorException;


class Device
{
    public $hostname, $net_address, $net_port;
    private $local_address, $MAC_address;
    private Network $network;
    private Router $router;

    public function __construct($net_address, $net_port)
    {
        $this->hostname = gethostbyaddr($net_address);
        $this->net_address = $net_address;
        $this->net_port = $net_port;
    }

    public function get_net_address()
    {
        return $this->net_address;
    }

    private function set_network(Network $network)
    {
        $this->network = $network;
    }

    public function set_local_address($local_address)
    {
        $this->local_address = $local_address;
    }

    public function get_local_address()
    {
        return $this->local_address;
    }

    public function set_MAC_address($MAC_address)
    {
        $this->MAC_address = $MAC_address;
    }

    public function find_network($searching_netmask = 32, $searching_step = 1)
    {

        while ($searching_netmask > 0) {

            $networks = DBAdapter::get_network_by_netmask($searching_netmask);

            $searching_netmask -= $searching_step;

            if (empty($networks))
                continue;


            foreach ($networks as $network) {
                if (IPConfig::ipv4_in_range($this->net_address, $network['network'])) {
                    $this->set_network(new Network($network));
                    return true;
                }
            }
        }

        return false;


    }

    public function gather_network_information()
    {
        if (!$this->find_network())
            throw new ErrorException("Network not found");

        if (!$this->network->find_routers())
            throw new ErrorException("Routers not found");

        $routers = $this->network->get_routers();

        foreach ($routers as $router) {
            if ($router->is_gateway())
                $this->router = $router;

            if ($router->get_address() == $this->net_address) {
                $this->router = $router;
                break;
            }

        }

        $this->router->get_remote_information($this);

        return [
            "hostname" => $this->hostname,
            "net_address" => $this->net_address,
            "net_port" => $this->net_port,
            "local_address" => $this->local_address,
            "MAC_address" => $this->MAC_address,
            "network" => $this->network->get_information(),
        ];


    }

}