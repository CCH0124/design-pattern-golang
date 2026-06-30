
舉例來說，在 Kubernetes 或是私有雲管理平台中，如果要擴容 (Scale out) 或 配置新的虛擬機器 (VM)，我們不可能每次都透過龐大的建構函式從零開始設定 CPU、記憶體、作業系統、基礎網路規則與預設的環境變數。相反地，我們會準備幾個標準的黃金映像檔 (Golden Image) 或樣板 (Template)，當需要新機器時，直接*複製 (Clone)*這個樣板，然後再微調特定的參數（例如：IP 位址、主機名稱）。

這正是 **原型模式 (Prototype Pattern)** 大顯身手的最佳生產環境場景。

---

1. 原型模式與泛型的完美結合

    **核心概念：** 定義一個能複製自身的介面，透過複製現有的「原型」物件來產生新物件，而不是實例化（`new`）一個新類別。

    **結合 Go 泛型的設計原則優勢：**
    1. **隱藏建立複雜物件的細節：** 客戶端只需要呼叫 `Clone()`，不需知道物件內部巢狀結構是如何被初始化的。
    2. **消滅型別轉型 (Type Casting)：** 傳統的原型模式 `Clone()` 通常回傳一個抽象的 `Prototype` 介面，客戶端拿到後還要強迫轉型為具體類別。透過 Go 泛型介面 `Cloneable[T any]`，複製出來的物件直接就是目標型別，乾淨且安全。
    3. **動態配置管理：** 在執行時期 (Runtime)，系統可以隨時註冊新的原型樣板到「原型註冊表 (Registry)」中，讓系統具備極高的擴充彈性。

2. 系統架構流程圖 (Class Diagram)

        我們將設計一個帶有泛型的原型註冊表架構，以下是生產環境級別的類別圖設計：

        ```plantuml
        @startuml
        skinparam classAttributeIconSize 0
        skinparam maxMessageSize 150

        ' Define the Generic Interface
        interface "Cloneable<T>" as Cloneable <<interface>> {
            + Clone(): T
        }

        ' Define the Concrete Prototype
            class VirtualMachine {
            - CPU: int
            - MemoryGB: int
            - OS: string
            - Tags: map[string]string
            + Clone(): *VirtualMachine
            + SetTag(key: string, value: string)
        }

        ' Define the Generic Registry
        class "PrototypeRegistry<T>" as PrototypeRegistry {
            - templates: map[string]T
            + Register(name: string, prototype: T)
            + Get(name: string): (T, error)
        }

        class Client {
        }

        ' Relationships
        VirtualMachine .up.|> Cloneable : " implements"
        PrototypeRegistry o-right-> Cloneable : " manages templates >"
        Client -down-> PrototypeRegistry : " looks up template"
        Client .right.> VirtualMachine : " clones & mutates"

        note bottom of VirtualMachine
            **Notice:**
            The **Clone()** method must implement a **Deep Copy**.
            Reference types like slices and maps (e.g., Tags) 
            must be re-allocated and copied manually to prevent 
            memory state pollution between instances.
        end note

        @enduml
        ```

3. 實戰演練題目：私有雲 VM 樣板配置系統

    **背景情境**
    身為基礎設施團隊的 SRE，你正在開發一套用 Go 撰寫的私有雲佈署工具。系統會預載幾種標準的 VM 樣板（例如：`web-server-base`, `db-server-base`）。
    由於 Go 語言中的 Map 與 Slice 是參考型別 (Reference Type)，如果只是單純賦值（Shallow Copy，淺拷貝），不同 VM 之間的標籤 (Tags) 將會互相污染，引發嚴重的資安與配置災難。

    **實作任務**
    請使用 Go 的泛型與原型模式，完成以下系統設計：

    1. **定義泛型介面**
    建立一個泛型介面 `Cloneable[T any]`，裡面包含一個 `Clone() T` 方法。

    2. **定義具體物件 (Concrete Prototype)**
        建立一個 `VirtualMachine` 結構，包含以下欄位：
        * `CPU` (int)
        * `MemoryGB` (int)
        * `OS` (string)
        * `Tags` (map[string]string) - *注意：這是坑點！*
        實作其 `Clone() *VirtualMachine` 方法，**務必確保實作了「深拷貝 (Deep Copy)」**，特別針對 `Tags`。

    3. **實作泛型原型註冊表 (Prototype Registry)：**
        建立一個泛型結構 `TemplateRegistry[T Cloneable[T]]`。
        包含兩個方法：
        * `Register(name string, tpl T)`：將樣板註冊到內部的 Map 中。
        * `Get(name string) (T, error)`：從 Map 中找出樣板並回傳它的 **Clone()**（而不是原本的樣板，以保護樣板不被竄改）。

    4. **撰寫 Client 端測試 (main 函式)：**
        * 實例化一個 `TemplateRegistry[*VirtualMachine]`。
        * 建立一個名為 `web-server` 的標準 VM 樣板 (2 CPU, 4GB, OS="Ubuntu", Tags={"role": "web"})，並註冊到 Registry。
        * 從 Registry `Get` 兩台機器出來 (VM1, VM2)。
        * 將 VM1 的 CPU 改為 4，並新增 tag `env=prod`。
        * 將 VM2 的 tag `env=dev` 新增進去。
        * 印出 原始樣板、VM1、VM2 的狀態，**驗證它們的 Tags 與數值完全獨立，沒有互相污染**。


在生產環境中，深拷貝 (Deep Copy) 是 Prototype Pattern 最核心的靈魂。請思考在 Go 中如何正確地走訪並複製一個 Map，而不是直接 `newMap = oldMap`。
