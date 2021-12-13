require 'rake/testtask'

Rake::TestTask.new(:test) do |t|
    # libs表示要添加到$LOAD_PATH(就是加载ruby文件的搜索路径)的文件夹
    # 默认是"lib"，现在再添加"test"
    t.libs << "algorithm"
    # 要运行的测试文件的特征。匹配以test_开头的所有文件
    t.pattern = 'algorithm/**/*.rb'
    # 不输出测试文件的信息
    t.verbose = false
end
  