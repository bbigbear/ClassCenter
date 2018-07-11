[<define "base_header">]
<div class="layui-header">
    <div class="layui-logo " style="color:#F0F0F0;font-size:20px;">课程管理中心</div>
    <ul class="layui-nav layui-layout-right">
      <li class="layui-nav-item">
        <a href="javascript:;">
          <img src="../static/img/admin.jpg" class="layui-nav-img">
          管理员
        </a>
        <dl class="layui-nav-child">
          <dd><a href="">基本资料</a></dd>
          <dd><a href="">安全设置</a></dd>
        </dl>
      </li>
      <li class="layui-nav-item"><a href="/">退出</a></li>
    </ul>
</div>
[<end>]

[<define "base_side">]
	<div class="layui-side-scroll">
      <ul class="layui-nav layui-nav-tree" lay-filter="test">
		<li class="layui-nav-item">
          <a class="" href="javascript:;">教学课程管理</a>
          <dl class="layui-nav-child">
            <dd><a href="javascript:;">教学课程申请</a></dd>			
            <dd><a href="javascript:;">教学课程审核</a></dd>			
            <dd><a href="javascript:;">教学课程维护</a></dd>					
          </dl>
        </li>
        <li class="layui-nav-item">
          <a class="" href="javascript:;">排课管理</a>
          <dl class="layui-nav-child">
            <dd><a href="javascript:;">选修课排课</a></dd>
			<dd><a href="javascript:;">必修课排课</a></dd>			
            <dd><a href="javascript:;">排课参数设置</a></dd>			
            <dd><a href="javascript:;">教师不排课时间设置</a></dd>			
            <dd><a href="javascript:;">教室不排课时间设置</a></dd>            			
          </dl>
        </li>
        <li class="layui-nav-item">
          <a href="javascript:;">上课时间设置</a>
          <dl class="layui-nav-child">
            <dd><a href="javascript:;">选修课课表设置</a></dd>
            <dd><a href="javascript:;">必修课课表设置</a></dd>
          </dl>
        </li>
        <li class="layui-nav-item">
		  <a href="javascript:;">教学计划管理</a>
		  <dl class="layui-nav-child">
            <dd><a href="javascript:;">计划申请</a></dd>
            <dd><a href="javascript:;">计划审核</a></dd>
            <dd><a href="javascript:;">计划维护</a></dd>
            <dd><a href="javascript:;">教学任务生成</a></dd>
			<dd><a href="javascript:;">教学任务分配</a></dd>
			<dd><a href="javascript:;">教学任务审核</a></dd>
          </dl>
		</li>
        <li class="layui-nav-item">
		  <a href="javascript:;">选修课管理</a>
		  <dl class="layui-nav-child">
			<dd><a href="javascript:;">选课设置</a></dd>
			<dd><a href="javascript:;">选课任务管理</a></dd>
			<dd><a href="javascript:;">开课申请</a></dd>
			<dd><a href="javascript:;">开课审核</a></dd>
          </dl>
		</li>
		<li class="layui-nav-item">
		  <a href="javascript:;">查询统计报表</a>
		  <dl class="layui-nav-child">
			<dd><a href="javascript:;">学生选课情况查询</a></dd>
          </dl>
		</li>
		<li class="layui-nav-item">
		  <a href="javascript:;">课表管理</a>
		  <dl class="layui-nav-child">
			<dd><a href="javascript:;">教师课表</a></dd>
			<dd><a href="javascript:;">教室课表</a></dd>
			<dd><a href="javascript:;">教学班课表</a></dd>
          </dl>
		</li>
		<li class="layui-nav-item">
		  <a href="javascript:;">协作组管理</a>
		  <dl class="layui-nav-child">
			<dd><a href="javascript:;">协作组维护</a></dd>
			<dd><a href="javascript:;">协作组讨论</a></dd>
          </dl>
		</li>
		<li class="layui-nav-item">
		  <a href="javascript:;">集备管理</a>
		  <dl class="layui-nav-child">
			<dd><a href="javascript:;">集备计划维护</a></dd>
			<dd><a href="javascript:;">待发起的集备</a></dd>
			<dd><a href="javascript:;">我发起的集备</a></dd>
			<dd><a href="javascript:;">我参加的集备</a></dd>
			<dd><a href="javascript:;">集备成功查看</a></dd>
          </dl>
		</li>
		<li class="layui-nav-item">
		  <a href="javascript:;">集备检查</a>
		  <dl class="layui-nav-child">
			<dd><a href="javascript:;">集备检查计划</a></dd>
			<dd><a href="javascript:;">我参加的检查</a></dd>
			<dd><a href="javascript:;">集备检查查看</a></dd>
          </dl>
		</li>
		<li class="layui-nav-item">
		  <a href="javascript:;">教师备课</a>
		  <dl class="layui-nav-child">
			<dd><a href="javascript:;">课时管理</a></dd>
			<dd><a href="javascript:;">我的教案</a></dd>
			<dd><a href="javascript:;">教研组维护</a></dd>
			<dd><a href="javascript:;">执行教案审核</a></dd>
			<dd><a href="javascript:;">入档案查询统计</a></dd>
			<dd><a href="javascript:;">教学进度统计报表</a></dd>
          </dl>
		</li>
      </ul>
    </div>
[<end>]