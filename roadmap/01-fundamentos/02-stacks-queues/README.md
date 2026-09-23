# Nivel 1.2: Pilas (Stacks - LIFO) y Colas (Queues - FIFO)

En Go **no existen clases ni tipos nativos `Stack` o `Queue`** en la biblioteca estándar como en Java (`java.util.Queue`) o Python (`collections.deque`). Los gophers construimos pilas y colas usando **slices**.

Pero acá hay dos peligros gravísimos de memoria:
1. **Memory Leaks por Reslicing**: Si hacés `q = q[1:]` para desencolar, los elementos viejos siguen existiendo en el array subyacente en la RAM, impidiendo que el Garbage Collector los limpie.
2. **LIFO vs FIFO**: Entender cuándo usar una Pila (deshacer/rehacer, parseo de sintaxis, llamada de funciones) y cuándo una Cola (buffers de mensajes, procesamiento de tareas en orden de llegada).
