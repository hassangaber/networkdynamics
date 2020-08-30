
<?php
//PHP file which finds the IP address of a remote user, even if the user is behind a proxy.

    function remoteIP(){

        $ipaddr = $_SERVER['$REMOTE_ADDR'];
        
        if (!empty($_SERVER['HTTP_CLIENT_IP'])) {

            $ipaddr = $_SERVER['HTTP_CLIENT_IP'];
        } 
        elseif (!empty($_SERVER['HTTP_X_FORWARDED_FOR'])) {

            $ipaddr = $_SERVER['HTTP_X_FORWARDED_FOR'];
        }
     
        return $ipaddr;
    }

    echo $ipaddr;
?>
