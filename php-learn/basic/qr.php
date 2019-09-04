<?php

include "./phpqrcode.php";

  $url="http://www.shiyanlou.com"; // 二维码url
  $path = false; // false 不生成图片文件
  $level = "L"; // 控制二维码容错率 默认L L M Q H
  $size = 10;// 生成图片的大小 默认为 4 1-10
  $margin = 2; //生成二维码的空白区域大小
  // saveandpoint # 保存二维码图片并显示出来
  $qrcode = new Qrcode(); 
  $qrcode->png($url, $path, $level, $size, $margin);
