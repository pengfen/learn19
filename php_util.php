<?php
namespace app\common\utils;

/**
 *  工具类
 *  父类中使用protected 实例化工具类
    protected $util;
    protected function initialize() {
    	$this->util = new Util();
    }
    子控制器中可以直接调用
    $this->util->getMonthStartTime();
 */

use Endroid\QrCode\QrCode;
use Endroid\QrCode\ErrorCorrectionLevel;
use think\facade\Env;

class Util
{
	/**
	 * 防表单重复提交
	 * @param $csrf_token
	 */
    // function checkrepet($csrf_token) {
    //     $auth = request()->header('Authorization');
    //     if (!$auth) $auth = $csrf_token;
    //     $redis = redis();
    //     $key = "repet_".$auth;
    //     $check = $redis->exists($key);  
    //     if ($check) { // 判断key是否存在
    //         $csrf = $redis->get($key);
    //         if ($csrf_token == $csrf) return 9005;
    //     }
    //     $redis->set($key, $csrf_token);  
    //     //过期时间为5秒   
    //     $redis->expire($key, 5);  
    // }

    /**
     * 获取月初始时间
     * @param $data array|int 参数集
     * @param $type int 时间类型 1 时间戳
     * @return $time int 月初始时间
     */
    public function getMonthStartTime($data, $type=0) {
    	if ($type) return strtotime(date("Y-m-1", $data));
        $time = time();
        $year = isset($data["year"])?$data['year']:date("Y", $time);
        $month = isset($data["month"])?$data['month']:date("m", $time);
        return strtotime($year."-".$month."-1");
    }

	/**
	 * 获取上个月初始时间
	 * @param $time 本月时间戳
	 * @return $time int 上月初始时间
	 */
	public function getYesMonthStartTime($time) {
		$year = date("Y", $time);
		$month = date("m", $time);

		if ($month == 1) {
			$month = 12;
			$year = $year - 1;
		} else {
			$month = $month - 1;
		}
		return strtotime($year."-".$month."-1");
	}

    /**
     * 计算月开始时间和结束时间
     * @param $time int 当前时间戳
     * @return $info array 时间结果集 ($start_time >= $time < $end_time)
     * @return $info['start_time'] int 月开始时间
     * @return $info['end_time'] int 月结果时间
     */
    public function getMonthTime($time) {
    	$month = date('m', $time);
        $year  = date('Y', $time);
        if ($month == 12) {
        	$last_month = 1;
        	$last_year = $year + 1;
        } else {
        	$last_month = $month + 1;
        	$last_year = $year;
        }

		$start_day = $year.'-'.$month.'-1'; // 如 2019/12/1 0:0:0
		$end_day = $last_year.'-'.$last_month.'-1'; // 2020/1/1 0:0:0
    	$info['start_time'] = strtotime($start_day);
    	$info['end_time'] = strtotime($end_day);
    	return $info;
    }

    /**
     * 计算年开始时间和结束时间
     * @param $time int 当前时间
     * @param $type int 时间类型　1 年 默认时间戳
     * @return $info array 时间结果集 ($start_time >= $time < $end_time)
     * @return $info['start_time'] 年开始时间
     * @return $info['end_time'] 年结束时间
     */
    public function getYearTime($time, $type=0) {
    	if ($type) {
    		$year = $time;
    	} else {
    		$year  = date('Y', $time);
    	}
        
        $last_year = $year + 1;

		$start_day = $year.'-1-1'; // 如 2019/1/1 0:0:0
		$end_day = $last_year.'-1-1'; // 2020/1/1 0:0:0
    	$info['start_time'] = strtotime($start_day);
    	$info['end_time'] = strtotime($end_day);
    	return $info;
    }

	/**
	 * 计算工龄 - 年
	 * @param time 时间
	 * @return $work_len 工龄
	 */
	public function computeWorkLength($time) {
		$work_len = (time() - $time) / (86400 * 365);
		return floor($work_len);
	}

	/**
	 * 计算工龄 - 日
	 * @param time 时间
	 * @return work_len 工龄(天)
	 */
	public function computeWorkDay($time) {
		$work_len = (time() - $time) / 86400;
		return floor($work_len);
	}

	/**
	 * 根据工龄计算起始时间
	 * @param $work_len 工龄
	 * @param start_time 起始时间
	 */
	public function computeStartTime($work_len) {
		return strtotime("-{$work_len} year");
	}

    /**
	 * 生成树结构 
	 * @param $data array 原始数据
	 * @param $pid int 父节点
	 * @param $parent_name string 父节点字段名
	 * @param $child_name string 子节点集字段名
	 * @param $curr_name string 当前节点字段名
	 * @return $tree array 返回树结果数据
	 */
	public function genTree($data, $pid=0, $parent_name='parent_id', $child_name='children', $curr_name='id') {
		$tree = '';
		foreach($data as $k => $v)
		{
			if($v[$parent_name] == $pid)
			{
				$v[$child_name] = $this->genTree($data, $v[$curr_name], $parent_name, $child_name, $curr_name); // 找子节点
				$tree[] = $v;
			}
		}
		return $tree;
	}

	/**
	 * 获取当前节点的所有父节点
	 * @param $data array 数据集
	 * @param $curr_node string 当前节点
	 * @param $curr_name string 当前节点字段名
	 * @param $parent_name string 父节点字段名
	 * @return $curr_node array 返回所有父节点
	 */
    public function get_parent_node($data, $curr_node, $curr_name='id', $parent_name='parent_id') {
        foreach($data as $k => $v) {
            if ($curr_node == 0) return $curr_node; // 顶级父节点 跨出循环
            if ($v[$curr_name] == $curr_node) {
                $curr_node .= '-'.$this->get_parent_node($data, $v[$parent_name], $curr_name, $parent_name);
                break;
            }
        }
        return $curr_node;
    }

    /**
     * 获取当前节点的所有子节点
     * @param $data array 数据集
     * @param $curr_node int 当前节点
     * @param $curr_name string 当前节点名
     * @param $parent_name string 当前父节点名
     * @return $tree array 返回子节点集
     */
	public function get_child_node($data, $curr_node, $curr_name='id', $parent_name='pid') {
		$tree = [];
		if ($data) {
			foreach($data as $k => $v) {
				if ($v[$parent_name] == $curr_node) {
					$ret = $this->get_child_node($data, $v[$curr_name]);
					if ($ret) {
						$tree[] = $v[$curr_name];
						$tree = array_merge($tree, $ret);
					} else {
						$tree[] = $v[$curr_name];
					}
				}
			}
		} 

		return $tree;
	}


	public function excel_index($start='A') {
		return array_merge(range($start, 'Z'), range('AA', 'AZ'), range('BA', 'BZ'));
	}


    /**
     * 生成二维码
     * @param $url      跳转链接
     * @param $logUrl    二维码中间logo
     * @param $showType  显示状态  1：直接显示 2 二维码图片地址
     * @return mixed
     * @throws \Endroid\QrCode\Exception\InvalidPathException
     */
    public function makeQrCode($url,$logUrl='',$showType=2) {

        $set_log =0;
        $qrCode = new QrCode($url);
        if($logUrl){  //设置二维码中间logo
            $qrCode->setLogoPath($logUrl);
            $qrCode->setLogoWidth(90);
        }
        $qrCode->setErrorCorrectionLevel(ErrorCorrectionLevel::HIGH);
        $name = rand(1,99999999).time();
        $dirPath  =  Env::get('root_path').'/public/static/qrcode/'.date("Ymd");
       if(!is_dir($dirPath)){
           mkdir($dirPath, 0777, true);
           chmod($dirPath, 0777);
       }
        $path ='./static/qrcode/'.date("Ymd")."/".$name.'.png';
        $qrCode->writeFile($path);
        //echo $qrCode->writeString();
        //@去操作这张图片
        //@删除文件
        //unlink($path);
        if($showType =="1"){
            // 展示二维码
            // 前端调用 例：<img src="http://localhost/qr.php?url=base64_url_string">
            header('Content-Type: ' . $this->_qr->getContentType());
            return $this->_qr->writeString();
        } else {
            //输出出url
             return "http://".$_SERVER['HTTP_HOST']."/$path";
        }

    }



}