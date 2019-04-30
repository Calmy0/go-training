package listpack

/* Список имеет указатель на "головной" элемент
 * каждый элемент имеет указатель на список и, стало быть, указатель на головной элемент.
 * Если prev или next элемента указывает на "головной элемент", то это значит - конец/начало списка
 */

type list struct {
	head *element
	len int
}

type element struct {
	value int
	next *element
	prev *element
	list *list
}

/* ***********************************************************************
 *			методы коллекции
 * *********************************************************************** */

/* Создает список
/* инициализирует "головной элемент" */
func New () (*list) {
	l := list{}
	e := new(element)
	l.head = e
	l.head.prev = e
	l.head.next = e
	l.head.value = -1

	l.len = 0;
	return &l
}

/* Length() – возвращает размер коллекции */
func (l *list) Length() int { return l.len; }

/* Add(value int) – добавляет элемент в коллекцию */
func (l *list) Add(value int) {
		// создаем элемент
	n := new(element)
	n.value = value
	n.list  = l
	l.len++

		// привязываем элемент
	n.prev	= l.head.prev
	n.next  = l.head

		// привязываемся к элементу
	l.head.prev.next = n
	l.head.prev      = n

	return
}

/* Last() – возвращает последний элемент в коллекции (значение элемента) */
func (l *list) Last() ( err, value int ) {
	if (l.len == 0) { return -1,0 }	 // нет элемента

	return 0, l.head.prev.value
}

/* First() – возвращает первый элемент в коллекции (значение элемента) */
func (l *list) First() ( err, value int ) {
	if (l.len == 0) { return -1,0 }	 // нет элемента

	return 0, l.head.next.value
}

/* Get(index int) – возвращает элемент с данным индексом из коллекции (первый элемент имеет index = 0) */
func (l *list) Get(index int) (err int, result *element) {
	if ( index > l.len) { return -1,nil }	 // нет элемента

	result = l.head.next;	// элемент с индексом 0
	for i:=0; i<index; i++ {
		result = result.next
	}

	return 0, result
}

/* Remove(index int) – удаляет элемент с данным индексом из коллекции */
func (l *list) Remove(index int) (err int) {
	if ( index > l.len) { return -1 }	 // нет элемента

	e := l.head.next;	// элемент с индексом 0
	for i:=0; i<index; i++ {
		e = e.next
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	e.list.len--
		// для чистки памяти
	e.prev = nil
	e.next = nil
	e.list = nil

	return 0
}

/* Print() – выводит коллекцию в консоль (в одну строку) */
func (l *list) Print() {
	if l.len == 0 { print("list is empty. \n"); }
	
	e:= l.head.next
	for i:=0; i<l.len; i++ {
		print(e.value," ")
		e = e.next
	}
}

/* ***********************************************************************
 *			методы узла
 * *********************************************************************** */

/* Value() – возвращает значение узла */
func (c *element) Value() (int) { return c.value; }

 /* Next() – возвращает следующий элемент коллекции 
 * указатель на "головной элемент" - это конец списка. вернуть nil
 */
 func (c *element) Next() ( err int, e *element) { 
	if (c.next != c.list.head) { return 0, c.next; }
	return -1, nil
}

/* Prev() – возвращает предыдущий элемент коллекции */
func (c *element) Prev() (err int, e *element) {
	if (c.prev != c.list.head) { return 0, c.prev; 	}
	return -1, nil
}

