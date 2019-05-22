package list

type element struct {
	value int
	next *element
	prev *element
}

/* ***********************************************************************
 *			методы узла
 * *********************************************************************** */

/* Value() – возвращает значение узла */
func (e *element) Value() (int) { return e.value; }

 /* Next() – возвращает следующий элемент коллекции  */
func (e *element) Next() (*element) { return e.next; }

/* Prev() – возвращает предыдущий элемент коллекции */
func (e *element) Prev() (*element) { return e.prev; }