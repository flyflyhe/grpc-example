<?php

require 'vendor/autoload.php';


$user = new \User\UserServiceClient('127.0.0.1:9002', ['credentials' => \Grpc\ChannelCredentials::createInsecure()]);
$req = new \User\UserRequest();
$req->setName("test");
[$res, $status] = $user->Search($req)->wait();
var_dump($res);
var_dump($res->getName());
var_dump($status);
