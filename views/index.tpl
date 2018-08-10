<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <meta name="apple-mobile-web-app-status-bar-style" content="black"> 
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="format-detection" content="telephone=no">
<title>课程管理中心</title>
<link rel="stylesheet" href="/static/css/layui.css">
<link rel="stylesheet" href="/static/css/bootstrap.min.css">
  <script src="//code.jquery.com/jquery-1.9.1.js"></script>
  <script src="//code.jquery.com/ui/1.10.4/jquery-ui.js"></script>
</head>
<body class="layui-layout-body">
<div class="layui-layout layui-layout-admin">
  <div class="layui-header">
    [<template "base_header".>]
  </div> 
  <div class="layui-side layui-bg-black">
    [<template "base_side".>]
  </div>
  <div class="layui-body">
    <!-- 内容主体区域 -->
    <div style="padding: 15px;">
		<div class="layui-tab" lay-filter="demo" lay-allowclose="true">
		  <ul class="layui-tab-title">
		    <li class="layui-this" lay-id="11">教学课程申请</li>
		  </ul>
		  <div class="layui-tab-content">
		    <div class="layui-tab-item layui-show"><iframe src='/v1/jxkc/apply' style="width:100%;height:800px;"></iframe></div>
		  </div>
		</div>
	</div>
  </div> 
  <div class="layui-footer">
    <!-- 底部固定区域 -->
    ©2018 智慧校园. All Rights Reserved
  </div>
</div>
<style>
	.layui-tab-title li:first-child > i {
		display: none;
		disabled:true
	}
</style>

<script src="/static/layui.js"></script>
<!--<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>-->

<script>
	//JavaScript代码区域
	layui.use(['element','layer','jquery','table'], function(){
	  var element = layui.element
		,form=layui.form
		,layer=layui.layer
		,$=layui.jquery
		,table=layui.table;
	  //layer.msg("你好");
	//自动加载
	$(function(){
		$( "#sortable" ).sortable();
	    $( "#sortable" ).disableSelection();			
	});
	var dic = {"教学课程申请":"/v1/jxkc/apply","教学课程审核": "/v1/jxkc/check","教学课程维护":"/v1/jxkc/maintain","计划申请":"/v1/jxjh/apply","计划审核": "/v1/jxjh/check","计划维护":"/v1/jxjh/maintain","教学任务生成":"/v1/jxrw/generate","教学任务分配":"/v1/jxrw/allot","教学任务审核":"/v1/jxrw/check",
	"排课参数设置":"/v1/pkgl/setting","教室不排课时间设置":"/v1/pkgl/classroom","教师不排课时间设置":"/v1/pkgl/teacher","选修课排课":"/v1/pkgl/eletive","必修课排课":"/v1/pkgl/compulsory","选修课课表设置":"/v1/sksj/eletive","必修课课表设置":"/v1/sksj/compulsory","选课设置":"/v1/xxkgl/setting",
	"选课任务管理":"/v1/xxkgl/task","开课申请":"/v1/xxkgl/apply","开课审核":"/v1/xxkgl/approve","教师课表":"/v1/kbgl/teacher","教室课表":"/v1/kbgl/classroom","教学班课表":"/v1/kbgl/teachingclass","课时管理":"/v1/jsbk/time_manage","我的教案":"/v1/jsbk/case","教研组维护":"/v1/jsbk/group",
	"执行教案审核":"/v1/jsbk/case/approve","入档案查询统计":"/v1/jsbk/case/query","教学进度统计报表":"/v1/jsbk/schedule","集备计划维护":"/v1/jbgl/plan","待发起的集备":"/v1/jbgl/plan/wait","我发起的集备":"/v1/jbgl/plan/launch","我参加的集备":"/v1/jbgl/plan/join","集备成功查看":"/v1/jbgl/plan/view",
	"集备检查计划":"/v1/jbjc/plan/launch","我参加的检查":"/v1/jbjc/plan/join","集备检查查看":"/v1/jbjc/plan/view","协作组维护":"/v1/xzz","协作组讨论":"/v1/xzz_discuss"};
	var noclickList=["教学课程管理","排课管理","上课时间设置","教学计划管理","选修课管理","查询统计报表","课表管理","协作组管理","集备管理","集备检查","教师备课"]
	var newarray=new Array()
	var list =[]
	list[0]="教学课程申请"
	element.on('nav(test)', function(elem){
	  //alert(dic[elem[0].textContent])
	  console.log(elem); //得到当前点击的DOM对象
	  console.log(elem[0].textContent);
	  console.log(list);
	  //console.log($.inArray(list, elem[0].textContent));
	  var i= $.inArray(elem[0].textContent,noclickList)
	  if (i==-1){
		ChangeTabs(elem[0].textContent) 
	  }	   	  	  
	});
	
	//切换tabs
	function ChangeTabs(n){
	  var index= $.inArray(n,list)
	  //var data=elem[0].textContent;	  
	  if(index==-1){
		//新增一个Tab项
		  list.push(n)
	      element.tabAdd('demo', {
	        title: n //用于演示
	        ,content: '<iframe src='+dic[n]+' style="width:100%;height:800px;"></iframe>'
	        ,id:list.length-1 //实际使用一般是规定好的id，这里以时间戳模拟下
	      });
		//切换
		  element.tabChange('demo', list.length-1);
	  }else{
		//切换到tab项
		  element.tabChange('demo', index);
	  }
	}
	
	//删除
	element.on('tabDelete(demo)', function(data){
	 // console.log(data); //当前Tab标题所在的原始DOM元素
	  console.log(data.index); //得到当前Tab的所在下标
	 // console.log(data.elem); //得到当前的Tab大容器
	  list.splice(data.index,1);
	});
	
	
			
  });
	
</script>

</body>
</html>