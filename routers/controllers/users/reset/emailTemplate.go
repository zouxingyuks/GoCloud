package reset

const defaultEmailTmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>重置密码</title>
</head>
<body style="background-color: #F3F4F6; padding: 5rem; display: flex; justify-content: center; align-items: center;">
<div style="background-color: #ffffff; padding: 2rem; border-radius: 0.5rem; box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); width: 100%; max-width: 400px;">
    <h1 style="font-size: 1.5rem; font-weight: 600; margin-bottom: 1rem; color: #2563EB;">重置您的 GoCloud 密码</h1>
    <p style="color: #4B5563; margin-bottom: 1rem;">您最近请求重置 GoCloud 账户密码。请点击下面的链接来设置新密码：</p>
    <a href="{{.ResetPasswordLink}}" target="_blank"
       style="display: inline-block; background-color: #2563EB; color: #ffffff; padding: 0.5rem 1rem; border-radius: 0.25rem; text-decoration: none;">点击这里设置新密码</a>
    <p style="color: #4B5563; margin-top: 1rem; margin-bottom: 0.5rem;">或者复制下面的链接到浏览器地址栏：</p>
    <p style="color: #4B5563; margin-bottom: 1rem; word-break: break-all;">{{.ResetPasswordLink}}</p>
    <p style="color: #4B5563; margin-bottom: 0.5rem;">如果您没有请求重置密码，请忽略此邮件，并确保您的账户安全。</p>
    <p style="color: #4B5563;">谢谢！</p>
    <p style="color: #4B5563; margin-top: 1rem;">GoCloud 团队</p>
</div>
</body>
</html>`
const Title = "GoCloud 密码重置"
