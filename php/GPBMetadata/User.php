<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

namespace GPBMetadata;

class User
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(hex2bin(
            "0a97010a0a757365722e70726f746f120475736572221b0a0b5573657252" .
            "657175657374120c0a046e616d65180120012809221c0a0c557365725265" .
            "73706f6e7365120c0a046e616d6518012001280932400a0b557365725365" .
            "727669636512310a0653656172636812112e757365722e55736572526571" .
            "756573741a122e757365722e55736572526573706f6e7365220062067072" .
            "6f746f33"
        ), true);

        static::$is_initialized = true;
    }
}
