<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>集备计划维护</title>
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
  <div class="layui-inline">
    <label class="layui-form-label">集备标题</label>
    <div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-inline">
    <label class="layui-form-label">日期</label>
    <div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
    </div>
	<div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-inline">
    <label class="layui-form-label">操作步骤</label>
    <div class="layui-input-inline" style="width: 150px;">
      <select name="Major" id="Major" lay-filter="status_select">
		    <option value="请选择" > 请选择</option>
			<option value="已创建" >已创建</option>
      </select>
    </div>
  </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-inline" style="padding-left:20px;">
    	<button class="layui-btn" id="query">查询</button>
		<button class="layui-btn" id="clear">清除条件</button>
		<button class="layui-btn" id="add">新建计划</button>
  	</div>
  </div>
</form>

<br><br>

	<table id="list" lay-filter="announcement" style="width:auto;"></table>
	<script type="text/html" id="barDemo">
		<a class="layui-btn layui-btn-normal layui-btn-xs" lay-event="edit">查看</a>
		<a class="layui-btn layui-btn-danger layui-btn-xs" lay-event="del">删除</a>
	</script>
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
	    ,url: '/v1/jbgl/plan/getdata'//数据接口
	    //,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[   
	       {field:'Master', title:'主备人', width:120}
		  ,{field:'Title',  title:'标题', width:120}
	      ,{field:'Explain',  title:'说明', width:120}
		  ,{field:'Founder',  title:'创建人', width:120}
		  ,{field:'CreateTime',  title:'创建日期', width:120}
		  ,{field:'Status',  title:'状态', width:120}
		  ,{fixed: 'right', title:'操作',width:120, align:'center', toolbar: '#barDemo'}
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
			  title: '查看课程',
			  //closeBtn: 0, //不显示关闭按钮
			  shadeClose: true,
			  shade: false,
			  area: ['893px', '600px'],
			 // offset: 'rb', //右下角弹出
			  //time: 2000, //2秒后自动关闭
			  maxmin: true,
			  anim: 2,
			  content: ['/v1/jbgl/plan/edit?id='+data.Id], //iframe的url，no代表不显示滚动条
			  cancel: function(index, layero){			  
				layer.close(index)
				window.location.reload();
			  	return false; 
			  },
		});
	    }else if(layEvent === 'del'){
		      layer.confirm('真的删除行么', function(index){
		        var jsData={'id':data.Id}
				$.post('/v1/jsbk/case/del', jsData, function (out) {
	                if (out.code == 200) {
	                    layer.alert('删除成功了', {icon: 1},function(index){
	                        layer.close(index);
	                        table.reload({});
	                    });
	                } else {
	                    layer.msg(out.message)
	                }
	            }, "json");
				obj.del(); //删除对应行（tr）的DOM结构
		        layer.close(index);
		        //向服务端发送删除指令
		      });
		    }
	  });
  
	$('#add').on('click',function(){
		layer.open({
			  type: 2,
			  title: '新建计划',
			  //closeBtn: 0, //不显示关闭按钮
			  shadeClose: true,
			  shade: false,
			  area: ['893px', '600px'],
			 // offset: 'rb', //右下角弹出
			  //time: 2000, //2秒后自动关闭
			  maxmin: true,
			  anim: 2,
			  content: ['/v1/jbgl/plan/add'], //iframe的url，no代表不显示滚动条
		});
		return false;
	});
	
	$('#query').on('click',function(){
		//alert("点击查询")
		var major=$("#Major").val()
		var grade=$("#PlanGrade").val()
		var classs=$("#PlanClass").val()
		var planId=$("#PlanId").val()
		table.reload('listReload',{
			where:{
				major:major,
				grade:grade,
				class:classs,
				planId:planId,
			}
		})
		return false;
	})
});
</script>

</body>
</html>