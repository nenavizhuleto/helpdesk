<?php

namespace App\Controllers;

use CodeIgniter\API\ResponseTrait;


class Helpdesk extends BaseController
{
    use ResponseTrait;
    public function getInfo()
    {

        $address = $_SERVER['REMOTE_ADDR'];
        $port = $_SERVER['REMOTE_PORT'];
        $host = gethostbyaddr($address);
        $db = \Config\Database::connect();
        $IPInvestigator = new \App\Helpers\IPInvestigator($db, $address, $port);

        $local_address = $IPInvestigator->local_address;
        $MAC = $IPInvestigator->MAC;


        $data = [
            "hostname" => $host ? $host : "undefined",
            "address" => $address,
            "local_address" => $local_address ? $local_address : "undefined",
            "port" => $port,
            "MAC" => $MAC ? $MAC : "undefined"
        ];

        return $this->respond($data);
    }


}