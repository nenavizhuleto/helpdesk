<?php
namespace App\Models;

use CodeIgniter\Model;



class ClientNetworkModel extends Model
{

    protected $table = 'client_networks';
    protected $primaryKey = 'id';
    protected $allowedFields = ['name', 'network', 'location', 'client_id'];


}