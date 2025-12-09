# Task Manager CLI

یه برنامه ساده و سبک برای مدیریت تسک‌ها مستقیماً تو ترمینال.

### ویژگی‌ها
- اضافه کردن تسک جدید
- نمایش لیست تسک‌ها (با وضعیت انجام‌شده یا نه)
- علامت زدن تسک به عنوان انجام‌شده
- حذف تسک
- پاک کردن همه تسک‌های انجام‌شده
- محیط تعاملی زیبا با رنگ و جدول مرتب
- ذخیره تسک‌ها تو فایل JSON تو پوشه خانگی

### نصب

go install github.com/mrnull2050/task-manager@latest


(یا اگه می‌خوای از سورس بسازی:
git clone https://github.com/mrnull2050/task-manager.git
cd task-manager
go build -o task-manager


و بعد فایل باینری رو به PATH اضافه کن.)

### استفاده
فقط بزن:

task-manager


وارد محیط تعاملی می‌شی. پرامپت چیزی شبیه اینه:

task> 


دستورات داخلش:
- `list` → نمایش همه تسک‌ها
- `add خرید نان و شیر` → اضافه کردن تسک جدید
- `done 2` → علامت زدن تسک شماره ۲ به عنوان انجام‌شده
- `delete 3` → حذف تسک شماره ۳
- `clear` → پاک کردن همه تسک‌های انجام‌شده
- `help` → نمایش راهنما
- `exit` یا `quit` → خروج

همین! ساده و سریع، بدون دردسر. �

![Screenshot 2025-12-09 172159](https://github.com/user-attachments/assets/7d17f14d-5012-4526-9bf2-fa9e4c49a021)

<img width="619" height="870" alt="image" src="https://github.com/user-attachments/assets/293265cc-4d3c-4fdc-99c2-d1f3c1917bc7" />



�
