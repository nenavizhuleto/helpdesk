<?php

namespace App\Controllers;

use CodeIgniter\API\ResponseTrait;


class Clients extends BaseController
{
    use ResponseTrait;

    public function index()
    {
        $clientsModel = model("App\Models\ClientModel");

        $response['clients'] = $clientsModel->findAll();

        return $this->respond($response);
    }

    public function show($id)
    {
        $clientsModel = model("App\Models\ClientModel");
        $response['client'] = $clientsModel->where('id', $id)->first();

        if (empty($response['client']))
            return $this->failNotFound();

        return $this->respond($response);
    }
}