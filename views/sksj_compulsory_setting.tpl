<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>必修课课表设置</title>
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
<form class="layui-form layui-form-pane1" action="">
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">上午课节数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="AmNum" autocomplete="off" class="layui-input">
	    </div>
		<div class="layui-form-mid layui-word-aux">*最小值为1</div>
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">下午课节数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="PmNum" autocomplete="off" class="layui-input">
	    </div>
		<div class="layui-form-mid layui-word-aux">*最小值为1</div>
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">晚上课节数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="NightNum" autocomplete="off" class="layui-input">
	    </div>
		<div class="layui-form-mid layui-word-aux">*最小值为1</div>
  </div>
  <div class="layui-form-item">
    <div class="layui-inline layui-layout-right" style="padding:10px;">
    	<button class="layui-btn" id="save">保存参数</button>
  	</div>
  </div>
<br>
  <table id="list" lay-filter="announcement" style="width:auto;"></table>
	<script type="text/html" id="barDemo">
		<a class="layui-btn layui-btn-normal layui-btn-xs" lay-event="edit">修改</a>		
	</script>
</form>
<script src="/static/layui.js"></script>
<!-- <script src="../build/lay/dest/layui.all.js"></script> -->

<script>
layui.use(['form','laydate','upload','jquery','layedit','element','table','laytpl'], function(){
  var form = layui.form
  ,laydate=layui.laydate
  ,upload = layui.upload
  , $ = layui.jquery
  ,layedit=layui.layedit
  ,element=layui.element
  ,table=layui.table
  ,laytpl = layui.laytpl;
	//自动加载
	$(function(){
				
	});	
 	//table 渲染
	  table.render({
	    elem: '#list'
	    ,height: 315
	    ,url: '/v1/sksj/getdata?style=1'//数据接口
	    //,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[   
	      //{field:'PlanId', title:'编号', width:120}
		  {field:'Name',  title:'名称', width:120}
	      ,{field:'StartTime',  title:'开始时间', width:120}
		  ,{field:'EndTime',  title:'结束时间', width:120}
		  ,{field:'Mon',  title:'周一', width:80}
		  ,{field:'Tue',  title:'周二', width:80}
		  ,{field:'Wed',  title:'周三', width:80}
		  ,{field:'Thu',  title:'周四', width:80}
		  ,{field:'Fri',  title:'周五', width:80}
		  ,{field:'Sat',  title:'周六', width:80}
		  ,{field:'Sun',  title:'周日', width:80}
		  ,{fixed: 'right', title:'操作',width:80, align:'center', toolbar: '#barDemo'}
	    ]]
	  });
	//监听工具条
		table.on('tool(announcement)', function(obj){ //注：tool是工具条事件名，test是table原始容器的属性 lay-filter="对应的值"
		    var data = obj.data //获得当前行数据
		    ,layEvent = obj.event; //获得 lay-event 对应的值
		    if(layEvent === 'edit'){
		      //layer.msg('查看操作');		
				  layer.open({
				  type: 2,
				  title: '修改课表',
				  //closeBtn: 0, //不显示关闭按钮
				  shadeClose: true,
				  shade: false,
				  area: ['893px', '600px'],
				 // offset: 'rb', //右下角弹出
				  //time: 2000, //2秒后自动关闭
				  maxmin: true,
				  anim: 2,
				  content: ['/v1/sksj/edit?id='+data.Id], //iframe的url，no代表不显示滚动条
				  cancel: function(index, layero){			  
					layer.close(index)
					window.location.reload();
				  	return false; 
				  },
				});
	    	}
	  });
	$('#save').on('click',function(){
		var data={
			'Id':1,
			'AmNum':parseInt($("#AmNum").val()),
			'PmNum':parseInt($("#PmNum").val()),
			'NightNum':parseInt($("#NightNum").val())
			};
				//发布
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/sksj/save",
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
		return false;
	})
});
</script>

</body>
</html>