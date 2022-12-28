<?php


namespace App\Libraries\Discovery;

class Network
{

    private $name, $network, $location;

    private $id, $client_id;

    private $routers;

    public function __construct($network)
    {
        $this->id = $network['id'];
        $this->name = $network['name'];
        $this->network = $network['network'];
        $this->location = $network['location'];
        $this->client_id = $network['client_id'];
    }

    public function find_routers()
    {
        $results = DBAdapter::get_network_routers_by_id($this->id);
        $this->routers = array_map(fn($el): Router => new Router($el), $results);
        return !empty($this->routers);
    }

    public function get_routers()
    {
        if (empty($this->routers))
            $this->find_routers();

        return $this->routers;
    }

    public function get_information()
    {
        return [
            "name" => $this->name,
            "network" => $this->network,
            "location" => $this->location
        ];
    }
}