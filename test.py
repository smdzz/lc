import os
import re

# 定义要处理的根目录
root_dir = 'simple'

# 遍历根目录下的所有目录
for folder_name in os.listdir(root_dir):
    folder_path = os.path.join(root_dir, folder_name)
    if os.path.isdir(folder_path):
        # 使用正则表达式提取目录名中的数字和汉字
        match = re.match(r'^(\d+\.)(.*)', folder_name)
        if match:
            numbers = match.group(1)
            chinese = match.group(2)

            # 构造新目录名
            new_folder_name = numbers.rstrip('.')
            new_folder_path = os.path.join(root_dir, new_folder_name)

            # 重命名目录
            print(folder_path, new_folder_path)
            os.rename(folder_path, new_folder_path)

            # 获取新目录下的唯一文件
            files = os.listdir(new_folder_path)
            if len(files) == 1:
                unique_file_path = os.path.join(new_folder_path, files[0])

                # 添加注释到唯一文件的开头
                with open(unique_file_path, 'r+', encoding='utf-8') as unique_file:
                    content = unique_file.read()
                    unique_file.seek(0, 0)
                    if chinese:
                        unique_file.write(f'// {chinese}\n\n' + content)

print("所有目录处理完毕。")