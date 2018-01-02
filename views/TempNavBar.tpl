{{define "navbar"}}
		<div class="navbar-header">
	      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
	        <span class="sr-only">Toggle navigation</span>
	        <span class="icon-bar"></span>
	        <span class="icon-bar"></span>
	        <span class="icon-bar"></span>
	      </button>
	      <a class="navbar-brand" href="#">课程中心</a>
	    </div>
        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav">
			<li class="dropdown">
	          <a href="#" class="dropdown-toggle" data-toggle="dropdown">选课 <span class="caret"></span></a>
	          <ul class="dropdown-menu">
	            <li><a href="/chooseClass" target="_blank">在线选课</a></li>
	            <li><a href="/queryClass" target="_blank">浏览课程</a></li>
	            <li><a href="#">修改选课</a></li>
	          </ul>
	        </li>
			<li class="dropdown">
	          <a href="#" class="dropdown-toggle" data-toggle="dropdown" >信息查询 <span class="caret"></span></a>
	          <ul class="dropdown-menu">
	            <li><a href="#">查询个人信息</a></li>
	            <li><a href="#">成绩查询</a></li>
	          </ul>
	        </li>
			<li class="dropdown">
	          <a href="#" class="dropdown-toggle" data-toggle="dropdown" >教学评价 <span class="caret"></span></a>
	          <ul class="dropdown-menu">
	            <li><a href="#">学生评价</a></li>
	          </ul>
	        </li>
          </ul>
          <ul class="nav navbar-nav navbar-right">
            <li><a href="#">退出</a></li>
          </ul>
        </div><!--/.nav-collapse -->
{{end}}