//js树
// <link href="/ztree/zTreeStyle.css" rel="stylesheet" type="text/css"/>
// <script type="text/javascript" src="/ztree/jquery.ztree.all.min.js"></script>
// <script type="text/javascript">
var ext = <?php echo json_encode($ext);?>;
var setting = {
    view: { // 处理是否显示图标
        showIcon: false
    },
    check: { // 处理是否显示选框等
        enable: true,  
        chkStyle: "checkbox",  
        chkboxType: { "Y": "ps", "N": "ps" }  
    }, 
    data: { // 加载数据
        simpleData: {
            enable: true
        }
    },
};
		
var zNodes = <?php echo json_encode($info)?>; // 获取树的相关数据
$(document).ready(function(){
	$.fn.zTree.init($("#treeDemo"), setting, zNodes); // 初始化
    $('#confirm').click(function(){
    var treeObj = $.fn.zTree.getZTreeObj("treeDemo"); // 获取树对象
    var nodes = treeObj.getCheckedNodes(true); // 获取选中节点
    for (var i = 0; i < nodes.length; i ++) {
    // 处理节点
    }
    $("#id_show_power").on("click",function(){ // 显示有权限部分
        var treeObj = $.fn.zTree.getZTreeObj('treeDemo');
        var nodes = treeObj.getCheckedNodes(false);
        //alert(nodes.length);
        for (var i = 0; i < nodes.length; i ++) {
            if (nodes[i].isParent) {
                var child = nodes[i].children;
                var flag = true;
                for (var j = 0; j < child.length; j ++) {
                    if (child[j].checked == true) {
                        flag = false;
                        break;
                    }
                    if (child[j].isParent && child[j].checked == false) {
                        var three = child[j].children;
                        for(var k = 0; k < three.length; k ++) {
                            if (three[k].checked == true) {
                                flag = false;
                                break;
                            }
                        }
                    }
                }
                if (flag == true) {
                    treeObj.hideNode(nodes[i]);
                }
            } else {
                treeObj.hideNode(nodes[i]);
            }        
        }
    });
    $("#id_show_all_power").on("click",function(){ // 显示所有权限
        var treeObj = $.fn.zTree.getZTreeObj('treeDemo');
        var nodes = treeObj.getNodesByParam("isHidden", true);
        treeObj.showNodes(nodes);
        treeObj.expandAll(true);
    })
    // 展开某一父节点下所有子节点
    var treeObj = $.fn.zTree.getZTreeObj("tree");
    var nodes = treeObj.getSelectedNodes();
    if (nodes.length>0) {
	    treeObj.expandNode(nodes[0], true, true, true);
    }
}