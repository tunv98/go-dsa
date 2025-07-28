# Array

+ cấu trúc dữ liệu tuyến tính
+ lưu trữ các phần tử cùng kiểu
+ kích thước cố định
+ trong bộ nhớ liên tiếp -> giúp CPU cache hiệu quả
+ truy cập theo index: O(1)
+ thêm, xóa O(n)
    + lí do: cần phải dịch toàn bộ các phần tử phía sau để tạo chỗ trống
    + thêm ở đầu mảng thì cũng vậy vì mảng là khối bộ nhớ liên tiếp
    + vì:
        + OS không cho bạn mở rộng mảng về phía trước
        + Máy tính không có chỗ trống trước địa chỉ 0x1000 (nơi bắt đầu mảng)
        + Mọi phần tử trong mảng phải giữ thứ tự liên tục trong bộ nhớ

| Kỹ thuật           | Mô tả                                                       |
|--------------------|-------------------------------------------------------------|
| Two pointers       | Duyệt từ 2 đầu mảng, áp dụng trong sắp xếp, tìm cặp giá trị |
| Sliding Window     | Tìm tổng lớn nhất trong cửa sổ dài k: O(n)                  |
| Prefix Sum         | Tiền xử lý mảng để truy vấn tổng đoạn nhanh hơn             |
| Difference Array   | Cập nhật mảng hàng loạt O(1)                                |
| Kadane’s Algorithm | Tìm dãy con có tổng lớn nhất                                |
| Binary Search      | Tìm kiếm nhanh trong mảng đã sắp xếp                        |
| Hash Map + Array   | Kết hợp để giải toán như Two Sum, Subarray Sum == K         |

+ Khác biệt giữa Static Array vs Dynamic Array

| Tiêu chí        | Static Array          | Dynamic Array           |
|-----------------|-----------------------|-------------------------|
| Kích thước      | Cố định khi khởi tạo  | Thay đổi khi chạy       |
| Cấp phát bộ nhớ | Tĩnh (trước khi chạy) | Động (khi cần)          |
| Cách truy cập   | Nhanh (O(1))          | Nhanh (O(1))            |
| Thêm, xóa       | Chậm (O(n))           | Nhanh (trung bình O(1)) |

//code

```go
package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[:] // slice dùng chung vùng nhớ
	s = append(s, 6)
	fmt.Println(s) // 1,2,3,4,5,6
}

``` 

+ Capacity (cap)
    + là số lượng phần tử tối đa mà slice có thể chứa mà không cần cấp phát lại bộ nhớ
    + nó có thể lớn hơn hoặc bằng độ dài (len) của slice
    + nếu bạn append vào slice mà vượt quá capacity, Go sẽ tự động cấp phát một mảng mới với kích thước lớn hơn và sao
      chép các phần tử cũ vào mảng mới

```go
package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:3] //2, 3 -> start index: 1, number of elements: 3

	fmt.Println(len(s)) // 2 (số phần tử hiện tại)
	fmt.Println(cap(s)) // 4 (từ arr[1] đến arr[4])
}
```

```text
type sliceHeader struct {
    ptr *ElementType // con trỏ tới phần tử đầu của slice trong mảng gốc
    len int          // số phần tử slice hiện có
    cap int          // dung lượng tối đa slice có thể mở rộng, tính từ ptr (thường khi gán thì sẽ lấy len của mảng gốc làm cap)
}
```

+ len(s) -> số phần tử hiện tại trong slice (20, 30)
+ cap(s) -> số phần tử tối đa mà slice có thể chứa mà không cần cấp phát lại bộ nhớ (20, 30, 40, 50)
  => cap(s) = số phần tử từ vị trí slice bắt đầu (ptr) đến cuối mảng gốc, vì slice là “cửa sổ” nhìn vào mảng gốc, không
  thể vượt ra ngoài vùng nhớ được cấp

---
Append

+ khi gọi append(slice, value)
+ nếu len(slice) < cap(slice)
    + thì giá trị sẽ được thêm vào cuối slice
    + len(slice) tăng lên 1
+ nếu len(slice) == cap(slice)
    + thì Go sẽ cấp phát một mảng mới với kích thước gấp đôi (hoặc lớn hơn) so với cap(slice)
    + sao chép các phần tử cũ từ slice vào mảng mới
    + sau đó thêm giá trị mới vào cuối mảng mới
    + len(slice) tăng lên 1, cap(slice) cũng tăng lên

? Tại sao việc giữ con trỏ &s[i] là nguy hiểm sau khi append(s,x)
+ Vì khi append, Go có thể cấp phát một mảng mới và sao chép các phần tử cũ vào mảng mới
+ Nếu bạn giữ con trỏ đến phần tử cũ, nó sẽ trỏ đến vùng nhớ đã bị giải phóng hoặc không còn hợp lệ nữa
---
Slice:

+ cấu trúc tham chiếu (reference type)
+ khi gán 1 slice cho biến khác, 2 biến này sẽ tham chiếu tới cùng vùng dữ liệu

```text
s1 := []int{1, 2, 3}
s2 := s1
s2[0] = 100
fmt.Println(s1) // [100 2 3] - s1 cũng thay đổi
```

+ Xóa phần tử trong slice

```text
s = append(s[:i], s[i+1:]...)
```

+ s[:i] -> lấy tất cả phần tử từ đầu đến i-1
+ s[i+1:] -> lấy tất cả phần tử từ i+1 đến hết
+ ... -> variadic expansion, giúp append có thể nhận nhiều phần tử từ slice

---

1. Khi tạo 1 slice nhỏ từ 1 slice lớn (hoặc mảng lớn), slice nhỏ vẫn giữ tham chiếu toàn bộ vùng nhớ ban đầu

```text
data := make([]byte, 1e6)   // slice lớn
small := data[:10]          // chỉ lấy 10 byte
// Nhưng small vẫn "giữ" toàn bộ vùng nhớ 1MB → GC không giải phóng được
```

Khắc phục

````text
tmp := make([]byte, len(small))
copy(tmp, small) // tmp dùng bộ nhớ mới, không giữ tham chiếu tới data
````

2. Cách clone 1 slice mà không giữ tham chiếu tới slice gốc

```text
original := []int{1, 2, 3, 4, 5}

#C1
clone := append([]int(nil), original...)

#C2
clone := make([]int, len(original))
copy(clone, original)
```

Slice nếu được truyền qua hàm hay gán = thì đều sẽ ko thay đổi được nội dung

3. Reuse slice to avoid GC

```text
buf := make([]byte, 0, 1024)
buf = buf[:0] // reset slice về độ dài 0 nhưng giữ nguyên capacity
```