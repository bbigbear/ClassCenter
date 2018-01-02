<!DOCTYPE html>
<html>
<head>
	{{template "header"}}
	<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
	<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
</head>
<body>

<div class="navbar navbar-default" >
	<div class="container">
		<div class="navbar-header">
		    <a class="navbar-brand" href="#">浏览课程</a>		
		</div>
	</div>
</div>
	<div class="container">	
		<div class="panel panel-default" style="height:600px">
		    <div class="panel-body">
				<form role="form">
		        <div class="form-group">
				<div class="row">
				<div class="col-sm-4">					
					<label for="name">学年</label>
					<select class="form-control">
						<option>2012-2013</option>
						<option>2013-2014</option>
						<option>2014-2015</option>
						<option>2015-2016</option>
						<option>2016-2017</option>
					</select>
				</div>
				<div class="col-sm-4">
					<label for="name">学期</label>
					<select class="form-control">
						<option>1</option>
						<option>2</option>
						<option>3</option>
					</select>
				</div>
				<div class="col-sm-3" style="padding-top:24px">				
					<button type="button" class="btn btn-primary">查询</button>
				</div>
				</div>
				</div>
				</form>
				<div class="table-responsive">
					<table class="table table-bordered">
						<caption>课表</caption>
						<thead>
							<tr >
								<th></th>
								<th >上午</th>
								<th>下午</th>
								<th>晚间</th>
							</tr>
						</thead>
						<tbody>
							<tr>
								<td>星期一</td>								
								<td>ios</td>
								<td>andorid</td>
								<td>python</td>
							</tr>
							<tr>
								<td>星期二</td>
								<td>goalng</td>
								<td>mysql</td>
								<td>sqlite</td>
							</tr>
							<tr>
								<td>星期三</td>
								<td></td>
								<td></td>
								<td></td>
							</tr>
							<tr>
								<td>星期四</td>
								<td></td>
								<td></td>
								<td></td>
							</tr>
							<tr>
								<td>星期五</td>
								<td></td>
								<td></td>
								<td></td>
							</tr>
							<tr>
								<td>星期六</td>
								<td></td>
								<td></td>
								<td></td>
							</tr>
							<tr>
								<td>星期七</td>
								<td></td>
								<td></td>
								<td></td>
							</tr>
						</tbody>
					</table>
				</div>  	
		    </div>
		</div>
	</div>

</body>
</html>
