from django.urls import path
urlpatterns=[
  path('admin/',admin.site.urls)
]

from django.urls import include,path
urlpatterns=[
  path('admin/',admin.site.urls)
  path('mysite/',include('mysite.urls'))
]


path()函数传递四个函数,两个必选:route 和 view;两个可选:kwargs 和 name
    route     一个包含URL模式的字符串.django处理请求时,从urlpatterns的第一个模式开始与请求的url比较,直到找到匹配的模式  
    view      django匹配到模式后,会以一个HttpRequest对象作为第一个参数,以route捕获的值作为关键字参数调用指定的视图函数(views.py)
    kwargs    传递给view的参数,以字典传递
    name      给url命名,便于在其他地方引用
