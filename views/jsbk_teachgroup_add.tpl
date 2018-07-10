<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>新增教研组</title>
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <meta name="apple-mobile-web-app-status-bar-style" content="black"> 
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="format-detection" content="telephone=no">
  <link rel="stylesheet" href="/static/css/layui.css">

<style>
body{padding: 10px;}
</style>
</head>
<body>
<form class="layui-form" action="">
  <div class="layui-form-item">
	<label class="layui-form-label">教研组名称</label>
	<div class="layui-inline">			    	
		<div class="layui-input-inline" style="width: 100px;">
      		<input type="text" name="GroupName" id="GroupName" autocomplete="off" class="layui-input">
    	</div>
	</div>	
  </div> 
  <div class="layui-form-item">
	  <div class="layui-inline">
	    <label class="layui-form-label">组长</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="GroupLeader" autocomplete="off" class="layui-input">
	    </div>
		<button class="layui-btn" id="choose">选择人员</button>
	  </div>		
  </div>
  <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">组员</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="GroupMember" autocomplete="off" class="layui-input">
	    </div>
		<button class="layui-btn" id="choose1">选择人员</button>
	  </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-inline layui-layout-right">
      <button class="layui-btn layui-btn-primary" id="add">保存</button>
    </div>
  </div>
</form>

<br><br><br>

<script src="http://code.jquery.com/jquery-2.1.1.min.js"></script>
<script src="/static/layui.js"></script>
<!-- <script src="../build/lay/dest/layui.all.js"></script> -->

<script>
layui.use(['form','laydate','upload','jquery','layedit','element'], function(){
  var form = layui.form
  ,laydate=layui.laydate
  ,upload = layui.upload
  , $ = layui.jquery
  ,layedit=layui.layedit
  ,element=layui.element;
	
	//文本域
	var index=layedit.build('detail',{
		hideTool:['image','face']
	});

	//数据上传
	function uploadData(){
		var data={
			'Name':$("#GroupName").val(),
			'GroupLeader':$("#GroupLeader").val(),
			'GroupMember':$("#GroupMember").val()
			};
			console.log(data)	
				//添加
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/jsbk/group/add_action",
					data:JSON.stringify(data),
					async:false,
					error:function(request){
						alert("post error")						
					},
					success:function(res){
						if(res.code==200){
							alert("保存成功")	
						}else{
							alert("保存失败")
						}						
					}
			  	});	
	}
	
	$('#add').on('click',function(){
		var len = $("#demoList tr").length
		if (len==0){
			uploadData()
		}
		return false;
	});
	
	$('#choose').on('click',function(){
		layer.open({
			type: 2,
			title: '组长选择',
			//closeBtn: 0, //不显示关闭按钮
			shadeClose: true,
			shade: false,
			area: ['700px', '500px'],
			anim: 2,
			content: ['/v1/jsbk/group/addleader'], //iframe的url，no代表不显示滚动条
			cancel: function(index, layero){			  
				layer.close(index)
				window.location.reload();
				return false; 
			},
		});				
		return false;
	});
	
	$('#choose1').on('click',function(){
		layer.open({
			type: 2,
			title: '组员选择',
			//closeBtn: 0, //不显示关闭按钮
			shadeClose: true,
			shade: false,
			area: ['700px', '500px'],
			anim: 2,
			content: ['/v1/jsbk/group/addpeople'], //iframe的url，no代表不显示滚动条
			cancel: function(index, layero){			  
				layer.close(index)
				window.location.reload();
				return false; 
			},
		});				
		return false;
	});
	
});
</script>

</body>
</html>