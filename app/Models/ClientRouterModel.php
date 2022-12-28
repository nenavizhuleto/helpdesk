<?php
namespace App\Models;

use CodeIgniter\Model;



class ClientRouterModel extends Model
{

    protected $table = 'client_routers';
    protected $primaryKey = 'id';
    protected $allowedFields = ['name', 'client_id', 'network_id'];


}