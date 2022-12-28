<?php
namespace App\Models;

use CodeIgniter\Model;



class RouterCredentialsModel extends Model
{

    protected $table = 'router_credentials';
    protected $primaryKey = 'id';
    protected $allowedFields = ['router_id', 'address', 'port', 'login', 'password', 'gateway'];


}