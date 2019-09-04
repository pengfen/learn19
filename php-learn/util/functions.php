<?php
class mysql {
    private $link;
    /**
     * 连接MYSQL函数
     * 自定义配置文件,包含需要使用的信息
     * @return resource
     */
    function connect ($host, $user, $pwd, $charset, $db_name){
        //连接mysql
        $this->link = @mysqli_connect($host, $user, $pwd) or die ('数据库连接失败<br/>ERROR '.mysqli_errno().':'.mysqli_error());
        //设置字符集
        mysqli_set_charset($this->link, $charset);
        //打开指定的数据库
        mysqli_select_db($this->link, $db_name) or die('指定的数据库打开失败');
        return $this->link;
    }

    /* array(
    'username'=>'cy',
    'password'=>'123456',
    'email'=>'cy@qq.com'
    ) */
    /**
     * 插入记录的操作
     * @param array $array
     * @param string $table
     * @return boolean
     */
    function insert($table, $array){
        $keys = join(',',array_keys($array)); // 获取数组中的键名
        $values = "'".join("','", array_values($array))."'"; // 获取数组中的值并使用单引号
        $sql = "insert {$table}({$keys}) VALUES ({$values})"; // insert wp_options(option_name,option_value,autoload) VALUES ('test_ricky','5','yes')
        $res = mysqli_query($this->link, $sql);
        if($res){
            return mysqli_insert_id($this->link);
        }else{
            return false;
        }
    }

    /**
     * MYSQL更新操作
     * @param array $array
     * @param string $table
     * @param string $where
     * @return number|boolean
     */
    function update($table, $array, $where=null){
        $sets = "";
        foreach ($array as $key=>$val){
            $sets .= $key."='".$val."',";
        }
        $sets = rtrim($sets,','); //去掉SQL里的最后一个逗号
        $where = $where==null?'':' WHERE '.$where;
        $sql = "UPDATE {$table} SET {$sets} {$where}";
        $res = mysqli_query($this->link, $sql);
        if ($res){
            return mysqli_affected_rows($this->link);
        }else {
            return false;
        }
    }

    /**
     * 删除记录的操作
     * @param string $table
     * @param string $where
     * @return number|boolean
     */
    function delete($table, $where=null){
        $where = $where==null?'':' WHERE '.$where;
        $sql = "DELETE FROM {$table}{$where}";
        $res = mysqli_query($this->link, $sql);
        if ($res){
            return mysqli_affected_rows($this->link);
        }else {
            return false;
        }
    }

    /**
     * 查询一条记录
     * @param string $sql
     * @param string $result_type MYSQLI_ASSOC 关联数组
     * @return boolean
     */
    function get_row($sql, $result_type=MYSQLI_ASSOC){
        $result = mysqli_query($this->link, $sql);
        if ($result && mysqli_num_rows($result) > 0){
            return mysqli_fetch_array($result, $result_type);
        }else {
            return false;
        }
    }

    /**
     * 查询一条记录
     * @param string $sql
     * @param string $result_type MYSQLI_NUM 索引数组
     * @return boolean
     */
    function get_var($sql, $result_type=MYSQLI_NUM){
        $result = mysqli_query($this->link, $sql);
        if ($result && mysqli_num_rows($result) > 0){
            $row =  mysqli_fetch_array($result, $result_type);
            return $row[0];
        }else {
            return false;
        }
    }

    /**
     * 得到表中的所有记录
     * @param string $sql
     * @param string $result_type
     * @return boolean
     */
    function get_result($sql,$result_type=MYSQLI_ASSOC){
        $result = mysqli_query($this->link, $sql);
        if ($result && mysqli_num_rows($result)>0){
            while ($row = mysqli_fetch_array($result,$result_type)){
                $rows[] = $row;
            }
            return $rows;
        }else {
            return false;
        }
    }

    /**
     * 取得结果集中的记录的条数
     * @param string $sql
     * @return number|boolean
     */
    function getTotalRows($sql){
        $result=mysqli_query($this->link, $sql);
        if($result){
            return mysqli_num_rows($result);
        }else {
            return false;
        }

    }

    /**
     * 释放结果集
     * @param resource $result
     * @return boolean
     */
    function  freeResult($result){
        return  mysqli_free_result($result);
    }

    /**
     * 断开MYSQL
     * @param resource $link
     * @return boolean
     */
    function close(){
        return mysqli_close($this->link);
    }

    /**
     * 得到客户端的信息
     * @return string
     */
    function getClintInfo(){
        return mysqli_get_client_info();
    }

    /**
     * 得到MYSQL服务器端的信息
     * @return string
     */
    function getServerInfo($link=null){
        return mysqli_get_server_info($link);
    }

    /**得到主机的信息
     * @return string
     */
    function getHostInfo($link=null){
        return mysqli_get_host_info($link);
    }

    /**
     * 得到协议信息
     * @return string
     */
    function getProtoInfo($link=null){
        return mysqli_get_proto_info($link);
    }
}