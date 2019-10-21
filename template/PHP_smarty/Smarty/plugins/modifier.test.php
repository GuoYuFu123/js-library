<?php
    function smarty_modifier_test($utime, $format) {
        return date($format, $utime);
    }
?>