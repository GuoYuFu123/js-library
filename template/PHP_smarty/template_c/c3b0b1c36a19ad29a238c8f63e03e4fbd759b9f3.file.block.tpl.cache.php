<?php /* Smarty version Smarty-3.1.16, created on 2019-10-21 08:01:08
         compiled from "tpl\block.tpl" */ ?>
<?php /*%%SmartyHeaderCode:228785dad6468a4a3a9-33780814%%*/if(!defined('SMARTY_DIR')) exit('no direct access allowed');
$_valid = $_smarty_tpl->decodeProperties(array (
  'file_dependency' => 
  array (
    'c3b0b1c36a19ad29a238c8f63e03e4fbd759b9f3' => 
    array (
      0 => 'tpl\\block.tpl',
      1 => 1571644866,
      2 => 'file',
    ),
  ),
  'nocache_hash' => '228785dad6468a4a3a9-33780814',
  'function' => 
  array (
  ),
  'cache_lifetime' => 120,
  'version' => 'Smarty-3.1.16',
  'unifunc' => 'content_5dad6468a79217_80688536',
  'variables' => 
  array (
    'str' => 0,
  ),
  'has_nocache_code' => false,
),false); /*/%%SmartyHeaderCode%%*/?>
<?php if ($_valid && !is_callable('content_5dad6468a79217_80688536')) {function content_5dad6468a79217_80688536($_smarty_tpl) {?><?php if (!is_callable('smarty_block_test2')) include 'D:\\wamp\\www\\test\\demo2\\Smarty\\plugins\\block.test2.php';
?><?php $_smarty_tpl->smarty->_tag_stack[] = array('test2', array('replace'=>"true",'maxnum'=>5)); $_block_repeat=true; echo smarty_block_test2(array('replace'=>"true",'maxnum'=>5), null, $_smarty_tpl, $_block_repeat);while ($_block_repeat) { ob_start();?>

<?php echo $_smarty_tpl->tpl_vars['str']->value;?>

<?php $_block_content = ob_get_clean(); $_block_repeat=false; echo smarty_block_test2(array('replace'=>"true",'maxnum'=>5), $_block_content, $_smarty_tpl, $_block_repeat);  } array_pop($_smarty_tpl->smarty->_tag_stack);?>
<?php }} ?>
