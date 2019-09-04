    // 生成对应url
    public function generate_url($url) {
        //$userig = "eJxlz8tOg0AUBuA9T0HY1uhcuGniwjSUoG2iFMS4mUyZKY60MB2GWmh8dytqxHi2-5fzn3M0TNO0kvnynOZ53Vaa6E5yy7wyLWCd-YZSCkaoJlixfyE-SKE4oWvN1RBCx3EQAGMjGK*0WItv4TrYu8Q29EekYSUZer522ADAk3H*EFEM4SJIp9HDtM1nwf3FjD5Xk0ZvkyILn3D3gg4xK8ve61H26q*6JboN-beoWABGRc1lusnjPr6JetzQ1e5R7SfKvSt3YbuZB1kaJR0A16NKLbb85ymIXM*2xwftuWpEXQ0AAXgSGHyOZbwbHwq9Xi8_";
        $userig = "eJxlj1FPgzAYRd-5FaSvGteWlTDfNkHFTIzZ2NheCLZl*4YrTVfm0PjfjagZiff1nNyb**G4rovm09lVwXndKJvbVkvkXrsIo8sz1BpEXtjcM*IflCcNRuZFaaXpIGGMUYz7DgipLJTwa7zCUYra8p5xEFXezfxUDDEmAaWjoK-ApoOP0eomfg4vyu2sqoz-ALrx3*8SNbnlcxUEK75rZbijYgMseeFA2Fu8HcdmOmnoUzrI9inc*2Q8yESxSIsoCdfRqCULDRmEzWnNlr1JC3v594l4dOgx2qNHaQ5Qq06gmDBCPfwd5Hw6X8JTXok_";
        $identifier = ""; 
        $sdkappid = "";
        $number = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
        $random = '';
        for($i = 0; $i < 32; $i ++) {
            $random .= $number[array_rand($number)];
        }
        return $url."?usersig=$userig&identifier=$identifier&sdkappid=$sdkappid&random=$random&contenttype=json";
    }

    // 加群(腾迅IM)
    public function joinim() {
        $url = $this->generate_url("https://console.tim.qq.com/v4/group_open_http_svc/add_group_member");
    }

    // 指定消息(腾迅IM)
    //public function 

    public function gener_random() {
        $number = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
        $number1 = [1, 2, 3, 4, 5, 6, 7, 8, 9];
        $random = '';
        for($i = 0; $i < 32; $i ++) {
            $random .= $number[array_rand($number)];
            if ($i == 0 && $random == 0) {
                $random = $number1[array_rand($number1)];
            }
        }
        return $random;
    }


    // 群发消息(腾迅IM)
    public function im() {
        $url = $this->generate_url("https://console.tim.qq.com/v4/group_open_http_svc/send_group_msg");

        $content = [
            "method" => "sendMessage",
            "data" => [
                "user" => ["nickName" => 'ricky'],
                "content" => "addddd",
            ]
        ];
        $content = json_encode($content);

        $data = [
            "GroupId" => "13074",
            "Random" => $this->gener_random(),
            //"From_Account" => "18516391218",
            "MsgBody" => [
                ["MsgType" => "TIMTextElem", "MsgContent" => ["Text" => $content]],
                //["MsgType" => "TIMTextElem", "MsgContent" => ["" => 6, "Data" => "abc\u0000\u0001"]],
            ],
            // "OfflinePushInfo" => [
            //     "PushFlag" => 0,
            //     // "Desc" => "离线推送内容",
            //     // "Ext" => "这是透传的内容",
            //     // "AndroidInfo" => ["Sound" => "android.mp3"],
            //     // "ApnsInfo" => ["Sound" => "apns.mp3", "BadgeMode" => 1, "Title" => "apns title", "SubTitle" => "apns subtitle", "Image" => "www.image.com"]
            // ]
        ];                                                                  
        $data_string = json_encode($data);  
        $ch = curl_init($url);        
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "POST");                            
        curl_setopt($ch, CURLOPT_POSTFIELDS, $data_string);  
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);  
        curl_setopt($ch, CURLOPT_HTTPHEADER, array(                   
            'Content-Type: application/json',  
            'Content-Length: ' . strlen($data_string))           
        );                                                                                                                     
        return curl_exec($ch); 
    }