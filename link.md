在HTML文档中，除了`<a>`标签之外，还有几个其他标签也可能包含链接：

1. `<link>`**标签**：用于定义与当前文档相关的资源，多用于引用CSS样式表。

   html

   ```
   <link rel="stylesheet" href="styles.css">
   ```

2. `<img>`**标签**：用于嵌入图像，可以包含一个链接，指向图像的源。

   html

   ```
   <img src="image.jpg" alt="Example Image">
   ```

3. `<script>`**标签**：用于定义客户端脚本，比如JavaScript代码，也可能包含一个链接，指向外部脚本文件。

   html

   ```
   <script src="script.js"></script>
   ```

4. `<iframe>`**标签**：用于在页面中嵌入另一个HTML页面。

   html

   ```
   <iframe src="https://example.com"></iframe>
   ```

5. `<form>`**标签**：虽然表单自身并不直接包含链接，但表单的`action`属性会指向一个URL，表示表单提交的目标地址。

   html

   ```
   <form action="submit-form.php" method="post">
       <input type="text" name="name">
       <input type="submit" value="Submit">
   </form>
   ```

6. `<source>`**标签**：用于多个元素，比如`<video>`和`<audio>`，指示媒体文件的来源。

   html

   ```
   <video controls>
       <source src="movie.mp4" type="video/mp4">
       <source src="movie.ogg" type="video/ogg">
       Your browser does not support the video tag.
   </video>
   ```

这些标签都可能包含链接或引用其他资源的URL。因此，在构建爬虫时，如果需要提取所有的外部链接，建议不仅仅检查`<a>`标签，还应考虑这些其他标签。希望这些信息对你有帮助！如果你有更多的问题或者需要更详细的解释，请告诉我。