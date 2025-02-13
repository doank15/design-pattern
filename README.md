# Software Design Pattern 
- It is a set of reused, and code-designed.
- It improves the reusability, readability, and reliability of code

# 1. Singleton pattern 
- It is **`a creational design pattern`** that limits the number of objects that can **initialize** of a class. Ensure that there is only one instance of a class in any case and provide global access point.
- This pattern is useful when you need to control global state or shared resource access.

# Characteristics
- There is only one instance object.
- You must create your own instance object.
- You must provide global access point to instance object.

# Advantages
- Ensure global uniqueness of resources or state within your app.
- Reduce system resource consumption and improve system efficiency.
# Disadvantages
- Difficult to test cuz singleton objects have the same lifecycle as the app, which make it difficult to isolate tests in unit tests.
# Application scenarios.
- Configuration manager: In an application, configuration information usually requires only one instance to be managed, which can ensure the consistency of configuration information.
- Connection pool: The database connection pool need to be limited the number of database connections để tránh việc quá nhiều kết nối và tiêu thụ tài nguyên quá mức.
- Logger: A logging system usually only needs one instance to ghi lại thông tin log của application để tránh việc lặp lại và confusion of log information.
- Application state management: In some apps, global state management, such as user session management or permission verification status.


# 2. Factory Pattern
- It is **`a creational design pattern`** that encapsulates the creation of an object, với subclasses deciding which class to instantiate. 
- This pattern make the code structure clearer and make it easy to replace or extend product classes.
# Characteristics
- Encapsulation: Encapsulates the creation process of the object in the factory class.
- Extensibility: Inheritance (kế thừa) and Polymorphism (đa hình) make it easy to add new product classes. 
- Abstraction: The factory method defines the interface for creating objects. But the creation of concrete objects is implemented by subclass.
-------------> Phương thức factory định nghĩa interface để tạo các đối tượng, nhưng việc tạo các đối tượng cụ thể được thực hiện bởi các lớp con.
# Advantages
- Phân chia việc khởi tạo và sử dụng các đối tượng giúp cải thiện tính độc lập giữa các modules.
- It's easy to expand and no need to modify existing code when adding new product categories ---> Phù hợp với Open/Closed Principle.
# Disadvantages
- Mỗi lần bổ sung product class thì yêu cầu bổ sung một factory class cụ thể, có thể dẫn tới việc tăng số lượng classes.
- The factory class centralizes the logic for all instances, which can lead to a factory class being too large.
# Application scenarios
- Database connection: Create database connection objects based on the different database types such as MySQL and Postgresql
- Payment gateway: Create corresponding payment processing for different payment methods such as credit card, paypal or wechat pay.
- Image processing: In image processing, image processor are created according to the different image types such as JPG, PNG.

# 3. Observe Pattern 
- It is **`a behavioral design pattern`** that defines a one-to-many dependency between objects. Do vậy, khi một object change state, all objects mà phụ thuộc vào nó sẽ được notified and automatically updated. 
- This pattern is well suited for implementing distributed event processing system.
# Characteristic
- One-to-may relationships: A topic can have **multiple** observes.
- Abstract coupling: There is an abstract coupling between the observer and the subject, and adding a new observer does not affect the existing system.
- Dynamic linkage: Observer can join or exist at any time.
# Advantages
- The coupling between objects is reduced, and the subject and observer are loosely coupled.
- Extensibility is good, and adding a new observer or topic class does not affect the existing class.
# Disadvantages
- When there are many observer objects, the distribution of notifications can pause performance issues.
- If the dependencies between the observer and the subject are too complex ---> It can make the system difficult to maintain. 
# Application scenarios
- Message system 
- Stock market: Khi stock price updated, tất cả các investor (observe) người subscribe sẽ được thông báo về latest price information.  
- Resource monitoring: Trong hệ thống giám sát, khi tài nguyên hệ thống (CPU, memory usage) đạt tới ngưỡng, hệ thống giám sát (observe) được thông báo và có hành động thích hợp.

# 4. Decorator Pattern
- It is a **`structural design pattern`** cho phép users có thể bổ sung thêm một cách linh hoạt các function or responsibilities vào một object bằng cách thêm một decorator object mà không phải modifying the original object.
# Characteristic
- Dynamic Scaling: Responsibilities có thể được thêm một cách linh hoạt at runtime.
- The decorator does not change the interface of the object.
- Flexibility: Multiple decorators can be combined to add multiple responsibilities to the object.
# Advantages
- The responsibilities of the added object are dynamic and revocable.
- You can wrap an object with multiple decorators and add multiple responsibilities
- Decorators and objects can be changed independently and are not coupled to each other.
# Disadvantage
- Sử dụng quá nhiều decorator làm cho hệ thống trở nên phức tạp và khó hiểu.
- Nó có thể gây ra multiple-decorator call --> ảnh hưởng tới hiệu suất.
# Application scenarios
- Logging: adding logging function without change original objects.
- Caching
- Resource management 
- Security control: authentication, permission checks
- Transaction processing

# 5. Strategy Pattern
- It is **`a behavioral design pattern**` định nghĩa một tập hợp các algorithms và đóng gói để chúng có thể hoán đổi cho nhau. 