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
		$ssh_connection = $this->create_ssh_connection($this->address, $this->port, $this->login, $this->password);
		$device->set_local_address($this->find_local_address($ssh_connection, $device));
		$device->set_MAC_address($this->find_MAC_address($ssh_connection, $device));
		ssh2_disconnect($ssh_connection);

	}

	public function get_address()
	{
		return $this->address;
	}

	public function is_gateway()
	{
		return $this->gateway;
	}

	private function exec_ssh_command($connection, $cmd)
	{

		if ($stream = ssh2_exec($connection, $cmd)) {
			stream_set_blocking($stream, true);

			$data = "";
			while ($buf = fread($stream, 4096)) {
				$data .= $buf;
			}
			fclose($stream);
			$data = explode("\n", $data);
			return $data;
		}
	}

	private function create_ssh_connection($address, $port, $login, $password)
	{
		$ssh_connection = ssh2_connect($address, $port);

		if (@ssh2_auth_password($ssh_connection, $login, $password)) {
			return $ssh_connection;
		}

		return false;
	}

	private function find_local_address($ssh_connection, Device $device)
	{
		$retry_count = 0;

		while ($retry_count <= 3) {
			$retry_count++;
			$cmd = "ip firewall connection print where dst-address=\"{$this->server_address}\"";
			$output = $this->exec_ssh_command($ssh_connection, $cmd);



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

	private function find_MAC_address($ssh_connection, Device $device)
	{

		$cmd = "ip arp print value-list where address=\"{$device->get_local_address()}\"";
		$output = $this->exec_ssh_command($ssh_connection, $cmd);

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