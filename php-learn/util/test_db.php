<?php

require_once "db.php";

$sql = "select ID, user_name from members";
$db = DB::getIntance();
