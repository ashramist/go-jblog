# 特效： 混合模式 blend modes 
  ## 渐变 linear-gradient(角度（angle），起始颜色(start-color)，终止颜色(end-color))
   ### background 属性说明
   background-image ：指定一个文件或者生成的颜色渐变作为背景图片
       除了可以接受图片Url 路径资源，还可以接收渐变函数
       定义一个从白色过渡到蓝色的渐变 
      background-iamge:linear-gradient(to right,white,blue);
   background-position:设置背景图片初始位置
   background-size:指定元素内背景图片的渲染尺寸
   background-repeat:决定在需要填充整个元素时，是否平铺图片
   background-origin:决定背景相对于元素的边框盒、内边距框盒（初始值）或内容盒子来定位
   background-clip: 指定背景是否应该填充边框盒（初始值）、内边距框盒或内容盒子
   background-attachment：指定背景图片是随着元素上下滚动（初始值）,还是固定在视口区域。注意，使用fixed值会对页面
                          性能产生负面影响
   background-color:指定纯色背景，渲染到背景图片下方。


  