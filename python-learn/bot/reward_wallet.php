<?php

// 请求数据
require_once "../util/config.php";
require_once "../util/functions.php";
$host = $config["host"];
$user = $config["user"];
$pass = $config["password"];
$charset = $config["charset"];
$db_name = $config["dbname"];

$wallet = $_GET["c"];
$user_id = $_GET["id"];

$db = new mysql();
$link = $db->connect($host, $user, $pass, $charset, $db_name);
$ret = "Sorry, your wallet address has been registered.";
// 判断钱包地址是否存在
$sql = "SELECT id FROM wallet where wallet = '{$wallet}' ";
$wallet_addr = $db->get_var($sql);
if (!$wallet_addr) {
    // 根据数据是否存在
    $sql = "SELECT id FROM wallet where telegram_id = $user_id";
    $id = $db->get_var($sql);
    $ret = "Sorry, you have already participated in this bonuty..";
    if (!$id) {
        $db->insert("wallet", [
            "wallet" => $wallet,
            "telegram_id" => $user_id,
            "create_time" => time()
        ]);
        $ret = "Congratulations! We've recorded your wallet address. You'll get 100 XXX when the bonuty ends , when we meet our followers Target ! ";
    }
}

$db->close();

echo json_encode(["msg" => $ret]);