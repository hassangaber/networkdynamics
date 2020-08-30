
<?php 
//PHP script which returns whether a domain inputed is true or false through port analysis.
     function checkDomain($domain,$server,$findText){ 

        $connection = fsockopen($server, 43); 
        if (!$connection)
             return false; 

        fputs($connection, $domain."\r\n"); 

         $responder =':'; 

        while(!feof($connection)) 
         $responder .= fgets($con,128); 
        
        fclose($con); 

        $conStatus = true;

        return strpos($response, $findText);
     }
        
        echo $conStatus;
        
?>
