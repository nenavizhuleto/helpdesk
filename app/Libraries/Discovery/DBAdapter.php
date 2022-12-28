<?php



namespace App\Libraries\Discovery;

class DBAdapter
{


    public static function get_network_by_netmask($netmask)
    {
        $db = \Config\Database::connect();

        $results = $db->query("
        SELECT * 
            FROM `client_networks` 
            WHERE `network` LIKE('%/{$netmask}')
        ")->getResultArray();

        return $results;
    }


    public static function get_network_routers_by_id($network_id)
    {
        $db = \Config\Database::connect();

        $results = $db->query("
        SELECT
        cr.name as name,
        rc.address as address,
        rc.port as port,
        rc.login as login,
        rc.password as password,
        rc.gateway as gateway
            FROM client_routers as cr
            INNER JOIN router_credentials as rc
            ON cr.id = rc.router_id
            WHERE cr.network_id = {$network_id}
        ")->getResultArray();

        return $results;
    }
}