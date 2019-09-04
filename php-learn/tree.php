<?php
// 文章分类
require "./config.php";
$host = $config["host"];
$user = $config["user"];
$password = $config["password"];
$dbname = $config["dbname"];
$conn = mysqli_connect($host, $user, $password);
mysqli_select_db($conn, $dbname);

// 1. 获取所有分类
$sql = "select t.term_id, t.`name`, tt.parent from terms t left JOIN term_taxonomy tt on t.term_id = tt.term_id WHERE taxonomy = 'category'";
$res = mysqli_query($conn, $sql);
$category = [];
if ($res) {
    while ($row = mysqli_fetch_assoc($res)) {
        $category[] = $row;
    }
}

mysqli_close($conn);

function getTree($data, $pId)
{
    $tree = '';
    foreach($data as $k => $v)
    {
        if($v['parent'] == $pId)
        { //父亲找到儿子
            $v['parent'] = getTree($data, $v['term_id']);
            $tree[] = $v;
        }
    }
    return $tree;
}

$tree = getTree($category, 0);
echo json_encode($tree);


// php 生成树结构
// 1.直接使用下标替换(id系列号必须是顺序且不断的)
$items = array(
    1 => array('id' => 1, 'pid' => 0, 'name' => '安徽省'),
    2 => array('id' => 2, 'pid' => 0, 'name' => '浙江省'),
    3 => array('id' => 3, 'pid' => 1, 'name' => '合肥市'),
    4 => array('id' => 4, 'pid' => 3, 'name' => '长丰县'),
    5 => array('id' => 5, 'pid' => 1, 'name' => '安庆市'),
);
print_r(generateTree($items));

function generateTree($items){
    foreach($items as $item)
    $items[$item['pid']]['son'][$item['id']] = &$items[$item['id']];
    return isset($items[0]['son']) ? $items[0]['son'] : array();
}

// 2.递归
$people = [];
foreach($users as $item) {
    $people[$item['parentid']][] = $item;
}
$info = $this->getTree($info, 0, $people);

function getTree($data, $pId, $users)
{
    $tree = '';
    foreach($data as $k => $v)
    {
        if($v['parentid'] == $pId)
        {
            $v['children'] = $this->getTree($data, $v['id'], $users); // 找子节点
            if (isset($users[$v['id']])) { // 处理组织下员工
            $v['users'] = $users[$v['id']];
            }
            $tree[] = $v;
        }
    }
    return $tree;
}

// 测试树列表
public function test_list() {
    $info = $this->t_company_wx_department->get_all_list();

    // 生成树节点
    $info = $this->genTree($info, 0);
    dd($info);

    // 获取某节点的所有父节点
    //$info = $this->get_parent_node($info,74);

    // 获取某节点的所有子节点
    $info = $this->get_child_node($info, 74);
    dd($info);
}

//
public function genTree($data,$pid) {
    $tree = '';
    foreach($data as $k => $v)
    {
        if($v['pId'] == $pid)
        {
        $v['children'] = $this->genTree($data, $v['id']); // 找子节点

        // if (isset($users[$v['id']])) {
        //     $v['users'] = $users[$v['id']];
        // }
        $tree[] = $v;
        }
    }
    return $tree;
}

/**
 * 获取某节点的所有父节点
 *
 * @param $data 数据列表
 * @param $parent 当前节点
 * @return array 获取parent节点的所有父节点
 */
public function get_parent_node($data, $parent) {
    foreach($data as $k => $v) {
    if ($parent == 0) {
            return $parent;
        }
        if ($v['id'] == $parent) {
            $parent .= '-'.$this->get_parent_node($data, $v['pId']);
            break;
        }
    }
    return $parent;
}

/**
 * 获取某节点的所有子节点
 * 
 * @param $data 数据列表
 * @param $child 当前节点
 * @return array 获取child节点的所有子节点
 */
public function get_child_node($data, $child) {
    $tree = [];
    foreach($data as $k => $v) {
        if ($v['pId'] == $child) {
            $ret = $this->get_child_node($data, $v['id']);
            // if (is_array($ret)) { // 数组方式显现
            //     $tree[] = $v;
            //     $tree = array_merge($tree, $ret);
            // } else {
            //     $tree[] = $v;
            // }
            if ($ret) {
                $tree[] = $v['id'];
                $tree = array_merge($tree, $ret);
            } else {
                $tree[] = $v['id'];
            }
        }
    }
    return $tree;
}