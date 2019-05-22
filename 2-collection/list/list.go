package list

// import "collection/element_pk"

type list struct {
	head *element
	len int
}

/* ***********************************************************************
 *			методы 
 * *********************************************************************** */

/* New() - Создает список */
func New () (*list) {
	l := list{}
	l.head = nil
	l.len = 0
	return &l
}

/* Length() – возвращает размер коллекции */
func (l *list) Length() int { return l.len; }

/* First() – возвращает первый элемент в коллекции */
func (l *list) First() ( *element ) { return l.head; }

/* Last() – возвращает последний элемент в коллекции */
func (l *list) Last() ( *element ) {
	e := l.First()
	if e == nil { return nil; }

	for ; e.next != nil; e = e.next {}
	return e
}

/* Add(value int) – добавляет элемент в коллекцию */
func (l *list) Add(value int) {
		// создаем элемент
	n := new(element)
	n.value = value
	n.next = nil

		// привязываем элемент
	n.prev      = l.Last()
	if n.prev == nil {	// if is 1st element
		l.head = n
	} else {
		n.prev.next = n
	}

	l.len++
	return
}

/* Get(index int) – возвращает элемент с данным индексом из коллекции (первый элемент имеет index = 0) */
func (l *list) Get(index int) (result *element, err int) {
	if ( index > l.len) { return nil, 1 }	 // нет элемента

	result = l.First();	// элемент с индексом 0
	for i:=0; i<index; i++ {
		result = result.next
	}

	return result, 0
}

/* Remove(index int) – удаляет элемент с данным индексом из коллекции */
func (l *list) Remove(index int) (err int) {
	if ( index > l.len) { return 1 }	 // нет элемента

	e := l.First();	// элемент с индексом 0
	for i:=0; i<index; i++ {
		e = e.next
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	l.len--

	return 0
}

/* Print() – выводит коллекцию в консоль (в одну строку) */
func (l *list) Print() {
	e:= l.First()
	if e == nil { 
		print("list is empty. \n")
	} else {
		for ; e != nil; {
			print(e.value," ")
			e = e.next
		}
		print("\n")
	}
	
}
