<?php /* Smarty version Smarty-3.1.16, created on 2019-10-21 07:36:44
         compiled from "tpl\test.tpl" */ ?>
<?php /*%%SmartyHeaderCode:177885dad12ea7402c2-27603881%%*/if(!defined('SMARTY_DIR')) exit('no direct access allowed');
$_valid = $_smarty_tpl->decodeProperties(array (
  'file_dependency' => 
  array (
    '993d0a8df898c0f4072435541bc301f02265491a' => 
    array (
      0 => 'tpl\\test.tpl',
      1 => 1571641609,
      2 => 'file',
    ),
  ),
  'nocache_hash' => '177885dad12ea7402c2-27603881',
  'function' => 
  array (
  ),
  'cache_lifetime' => 120,
  'version' => 'Smarty-3.1.16',
  'unifunc' => 'content_5dad12ea740ae4_92573609',
  'variables' => 
  array (
    'articletitle' => 0,
    'arr' => 0,
    'time' => 0,
    'url' => 0,
    'article' => 0,
    'ifkey' => 0,
    'score' => 0,
    'list' => 0,
    'item' => 0,
  ),
  'has_nocache_code' => false,
),false); /*/%%SmartyHeaderCode%%*/?>
<?php if ($_valid && !is_callable('content_5dad12ea740ae4_92573609')) {function content_5dad12ea740ae4_92573609($_smarty_tpl) {?><?php if (!is_callable('smarty_modifier_capitalize')) include 'D:\\wamp\\www\\test\\demo2\\Smarty\\plugins\\modifier.capitalize.php';
if (!is_callable('smarty_modifier_date_format')) include 'D:\\wamp\\www\\test\\demo2\\Smarty\\plugins\\modifier.date_format.php';
?><?php echo smarty_modifier_capitalize($_smarty_tpl->tpl_vars['articletitle']->value);?>

<?php echo ($_smarty_tpl->tpl_vars['articletitle']->value).(" yesterday").(" sdfsf");?>


<br />
<?php echo $_smarty_tpl->tpl_vars['arr']->value['title']['title'];?>

<?php echo $_smarty_tpl->tpl_vars['arr']->value['author'];?>


<br />
<?php echo $_smarty_tpl->tpl_vars['arr']->value["title"]['title'];?>

<?php echo $_smarty_tpl->tpl_vars['arr']->value['author'];?>


<br />
<?php echo smarty_modifier_date_format($_smarty_tpl->tpl_vars['time']->value,"%Y/%m/%d %H:%M:%S %B %e");?>


<br />
<?php echo mb_strtoupper($_smarty_tpl->tpl_vars['url']->value, 'UTF-8');?>


<br />
<?php echo nl2br($_smarty_tpl->tpl_vars['article']->value);?>


<br />
<?php if ($_smarty_tpl->tpl_vars['ifkey']->value=="guoguo") {?>
guoguo
<?php } elseif ($_smarty_tpl->tpl_vars['ifkey']->value=="haha") {?>
haha
<?php } else { ?>
hehe
<?php }?>

<br />
<?php if ($_smarty_tpl->tpl_vars['score']->value>90) {?> 
90+
<?php } elseif ($_smarty_tpl->tpl_vars['score']->value>60) {?> 
60+
<?php } else { ?>
0+
<?php }?>

<br />
<?php if (isset($_smarty_tpl->tpl_vars['smarty']->value['section']['list'])) unset($_smarty_tpl->tpl_vars['smarty']->value['section']['list']);
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['name'] = 'list';
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['loop'] = is_array($_loop=$_smarty_tpl->tpl_vars['list']->value) ? count($_loop) : max(0, (int) $_loop); unset($_loop);
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['max'] = (int) 1;
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['show'] = true;
if ($_smarty_tpl->tpl_vars['smarty']->value['section']['list']['max'] < 0)
    $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['max'] = $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['loop'];
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['step'] = 1;
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['start'] = $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['step'] > 0 ? 0 : $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['loop']-1;
if ($_smarty_tpl->tpl_vars['smarty']->value['section']['list']['show']) {
    $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['total'] = min(ceil(($_smarty_tpl->tpl_vars['smarty']->value['section']['list']['step'] > 0 ? $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['loop'] - $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['start'] : $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['start']+1)/abs($_smarty_tpl->tpl_vars['smarty']->value['section']['list']['step'])), $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['max']);
    if ($_smarty_tpl->tpl_vars['smarty']->value['section']['list']['total'] == 0)
        $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['show'] = false;
} else
    $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['total'] = 0;
if ($_smarty_tpl->tpl_vars['smarty']->value['section']['list']['show']):

            for ($_smarty_tpl->tpl_vars['smarty']->value['section']['list']['index'] = $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['start'], $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['iteration'] = 1;
                 $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['iteration'] <= $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['total'];
                 $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['index'] += $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['step'], $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['iteration']++):
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['rownum'] = $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['iteration'];
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['index_prev'] = $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['index'] - $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['step'];
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['index_next'] = $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['index'] + $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['step'];
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['first']      = ($_smarty_tpl->tpl_vars['smarty']->value['section']['list']['iteration'] == 1);
$_smarty_tpl->tpl_vars['smarty']->value['section']['list']['last']       = ($_smarty_tpl->tpl_vars['smarty']->value['section']['list']['iteration'] == $_smarty_tpl->tpl_vars['smarty']->value['section']['list']['total']);
?>
    <?php echo $_smarty_tpl->tpl_vars['list']->value[$_smarty_tpl->getVariable('smarty')->value['section']['list']['index']]['title'];?>

    <br />
<?php endfor; endif; ?>

<br />

<?php  $_smarty_tpl->tpl_vars['item'] = new Smarty_Variable; $_smarty_tpl->tpl_vars['item']->_loop = false;
 $_from = $_smarty_tpl->tpl_vars['list']->value; if (!is_array($_from) && !is_object($_from)) { settype($_from, 'array');}
foreach ($_from as $_smarty_tpl->tpl_vars['item']->key => $_smarty_tpl->tpl_vars['item']->value) {
$_smarty_tpl->tpl_vars['item']->_loop = true;
?>
    <?php echo $_smarty_tpl->tpl_vars['item']->value["age"];?>


<?php }
if (!$_smarty_tpl->tpl_vars['item']->_loop) {
?>
当前没有文章
<?php } ?>

<br />
<?php  $_smarty_tpl->tpl_vars['item'] = new Smarty_Variable; $_smarty_tpl->tpl_vars['item']->_loop = false;
 $_from = $_smarty_tpl->tpl_vars['list']->value; if (!is_array($_from) && !is_object($_from)) { settype($_from, 'array');}
foreach ($_from as $_smarty_tpl->tpl_vars['item']->key => $_smarty_tpl->tpl_vars['item']->value) {
$_smarty_tpl->tpl_vars['item']->_loop = true;
?>
    <?php echo $_smarty_tpl->tpl_vars['item']->value['author'];?>

<?php } ?>



<?php echo $_smarty_tpl->getSubTemplate ("header.tpl", $_smarty_tpl->cache_id, $_smarty_tpl->compile_id, 9999, null, array('title'=>"网站头部",'table_bgcolor'=>"#e6e6e6",'sitename'=>"果果网"), 0);?>


<br />

<?php echo date('Y-m-d',$_smarty_tpl->tpl_vars['time']->value);?>

<br />


<?php echo str_replace("q","Q",$_smarty_tpl->tpl_vars['articletitle']->value);?>


<br />

<?php echo test(array('p1'=>"guoguo",'p2'=>"最帅"),$_smarty_tpl);?>
 
<?php }} ?>
