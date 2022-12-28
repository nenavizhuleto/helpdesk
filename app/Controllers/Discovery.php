<?php

namespace App\Controllers;

use App\Libraries\Discovery\Device;
use CodeIgniter\API\ResponseTrait;


class Discovery extends BaseController
{

    use ResponseTrait;

    public function index()
    {
        $address = $_SERVER['REMOTE_ADDR'];
        $port = $_SERVER['REMOTE_PORT'];
        try {
            $device = new Device($address, $port);

            // var_dump($address);
            // var_dump($port);

            $data = $device->gather_network_information();

            return $this->respond($data, 200);
        } catch (\Throwable $th) {
            return $this->respond([
                "message" => $th->getMessage(),
                "error" => $th->__toString()
            ], 300);
        }




    }
}