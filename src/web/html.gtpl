<html>
<head>
<title></title>
</head>
<body>
<!--enctype="multipart/form-data"不能取到参数-->
<form action="/html" method="post">
    用户名：<input type="text" name="username">
    <br>
    年龄：<input type="text" name="age">
    <br>
    email：<input type="text" name="email">
    <br>
    手机号：<input type="text" name="mobile">
    <br>
    水果：<select name="fruit">
       <option value="apple">apple</option>
       <option value="pear">pear</option>
       <option value="banane">banane</option>
       </select>
    <br>
    性别：<input type="radio" name="gender" value="1">男
    <input type="radio" name="gender" value="2">女
    <br>
    兴趣：<input type="checkbox" name="interest" value="football">足球
    <input type="checkbox" name="interest" value="basketball">篮球
    <input type="checkbox" name="interest" value="tennis">网球
    <br>
    照片：<input type="file" name="uploadfile" />

    <br><br>
    <input type="submit" value="登陆">
</form>
</body>
</html>