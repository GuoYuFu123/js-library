{$articletitle|capitalize}
{$articletitle|cat:" yesterday":" sdfsf"}

<br />
{$arr.title.title}
{$arr.author}

<br />
{$arr["title"]['title']}
{$arr['author']}

<br />
{$time|date_format:"%Y/%m/%d %H:%M:%S %B %e"}

<br />
{$url|upper}

<br />
{$article|nl2br}

<br />
{if $ifkey eq "guoguo"}
guoguo
{elseif $ifkey eq "haha"}
haha
{else}
hehe
{/if}

<br />
{if $score gt 90} 
90+
{elseif $score gt 60} 
60+
{else}
0+
{/if}

<br />
{section name=list loop=$list max=1}
    {$list[list].title}
    <br />
{/section}

<br />

{foreach from=$list item=item}
    {$item["age"]}

{foreachelse}
当前没有文章
{/foreach}

<br />
{foreach $list as $item}
    {$item.author}
{/foreach}

{* 文件引用 *}

{include file="header.tpl" title="网站头部" table_bgcolor="#e6e6e6" sitename="果果网"}

<br />
{* date("Y-m-d", $time) *}
{'Y-m-d'|date: $time}
<br />

{* str_replace("q", "Q", $articletitle) *}
{"q"|str_replace:"Q": $articletitle}

<br />
{* 调用f_test函数，属性会整合为参数数组即array("p1"=>"guoguo","p2"=>"最帅"...) *}
{f_test p1="guoguo" p2="最帅"} 
