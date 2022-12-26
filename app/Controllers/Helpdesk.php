<?php

namespace App\Controllers;

class Helpdesk extends BaseController
{
    public function index()
    {
        $address = $_SERVER['REMOTE_ADDR'];
        $port = $_SERVER['REMOTE_PORT'];
        $host = gethostbyaddr($address);
        $MAC = "undefined";


        $network = new \App\Models\NetworkModel();

        $data = [
            "networks" => $network->findAll(),
            "hostname" => $host ? $host : "undefined",
            "address" => $address,
            "port" => $port,
            "mac" => $MAC
        ];

        return view('helpdesk', $data);
    }
}