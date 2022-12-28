<?php

namespace App\Libraries\Discovery;


class Router
{
	private $name, $address, $port, $login, $password, $gateway;

	private $server_address;

	public function __construct($router)
	{
		$this->name = $router['name'];
		$this->address = $router['address'];
		$this->port = $router['port'];
		$this->login = $router['login'];
		$this->password = $router['password'];
		$this->gateway = $router['gateway'];

		$this->server_address = "{$_SERVER['SERVER_NAME']}:{$_SERVER['SERVER_PORT']}";
	}


	public function get_remote_information(Device $device)
	{
		$device->set_local_address($this->find_local_address($device));
		$device->set_MAC_address($this->find_MAC_address($device));

	}

	public function get_address()
	{
		return $this->address;
	}

	public function is_gateway()
	{
		return $this->gateway;
	}

	private function exec_ssh_command($cmd)
	{
		$ssh_cmd = "ssh -l {$this->login} -p {$this->port} {$this->address} '{$cmd}'";
		exec($ssh_cmd, $output);

		return $output;
	}

	private function find_local_address(Device $device)
	{
		$retry_count = 0;

		while ($retry_count <= 3) {
			$retry_count++;
			$cmd = "ip firewall connection print where dst-address=\"{$this->server_address}\"";
			$output = $this->exec_ssh_command($cmd);

			if (empty($output)) {
				continue;
			}

			foreach ($output as $line) {
				if (str_contains($line, "{$this->server_address}") && str_contains($line, "{$device->net_port}")) {
					return IPConfig::str_extractIPv4WithPort($line, $device->net_port);
				}
			}
		}

		throw new \ErrorException("Exceeded maximum retry amount while getting local ip address of {$device->net_address}");

	}

	private function find_MAC_address($device)
	{

		$cmd = "ip arp print value-list where address=\"{$device->get_local_address()}\"";
		$output = $this->exec_ssh_command($cmd);

		if (empty($output)) {
			throw new \ErrorException("Failed getting MAC address of {$device->net_address}");
		}

		foreach ($output as $line) {
			if (str_contains($line, "mac-address")) {
				$MAC = trim(explode(":", $line, 2)[1]);
			}
		}

		return $MAC;
	}


}