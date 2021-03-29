<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
</head>
<body>
  <?php
  $x = 'zhangyangyang';//
 $cars=array("Volvo","BMW","Toyota");
  echo "first php";
  function test (){
      global $x,$cars;
      echo $x;
      echo $cars;
      echo "<br>";
      var_dump($cars);
  }
  test();
  function mysql(){
    $servername = "localhost";
$username = "root";
$password = "root";
 
// 创建连接
$conn = new mysqli($servername, $username, $password);
 
// 检测连接
if ($conn->connect_error) {
    die("连接失败: " . $conn->connect_error);
} 
echo "连接成功";
  }
 // mysql();
  //phpinfo();
  echo $_SERVER['SERVER_ADDR'];
  echo $_SERVER['PHP_SELF'];
  ?>

<form method="post" action="<?php echo $_SERVER['PHP_SELF'];?>">
Name: <input type="text" name="fname">
<input type="submit">
</form>
 
<?php 
$name = $_POST['fname']; 
echo $name; 
?>
</body>
</html>