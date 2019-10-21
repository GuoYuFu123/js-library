<?php /* Smarty version Smarty-3.1.16, created on 2019-10-21 07:44:30
         compiled from "tpl\time.tpl" */ ?>
<?php /*%%SmartyHeaderCode:104975dad614b8b64f8-16287825%%*/if(!defined('SMARTY_DIR')) exit('no direct access allowed');
$_valid = $_smarty_tpl->decodeProperties(array (
  'file_dependency' => 
  array (
    '17e0c84a01ee004b4810c5fb83e8fe0040ddb1b0' => 
    array (
      0 => 'tpl\\time.tpl',
      1 => 1571643840,
      2 => 'file',
    ),
  ),
  'nocache_hash' => '104975dad614b8b64f8-16287825',
  'function' => 
  array (
  ),
  'cache_lifetime' => 120,
  'version' => 'Smarty-3.1.16',
  'unifunc' => 'content_5dad614b9135f0_69465374',
  'variables' => 
  array (
    'time' => 0,
  ),
  'has_nocache_code' => false,
),false); /*/%%SmartyHeaderCode%%*/?>
<?php if ($_valid && !is_callable('content_5dad614b9135f0_69465374')) {function content_5dad614b9135f0_69465374($_smarty_tpl) {?><?php if (!is_callable('smarty_modifier_test')) include 'D:\\wamp\\www\\test\\demo2\\Smarty\\plugins\\modifier.test.php';
?><?php echo $_smarty_tpl->tpl_vars['time']->value;?>

<?php echo smarty_modifier_test($_smarty_tpl->tpl_vars['time']->value,"Y-m-d");?>
<?php }} ?>
