<?php

namespace App\Controllers;

use CodeIgniter\API\ResponseTrait;

class Index extends BaseController
{
    use ResponseTrait;

    public function getIndex()
    {
        return $this->respond(["a" => 5]);
    }
}