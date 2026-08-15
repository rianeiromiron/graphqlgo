# Manual de Usuario - GraphQL Go CRUD

## Tabla de Contenidos
1. [Introducción](#introducción)
2. [Primeros Pasos](#primeros-pasos)
3. [Interfaz de Usuario](#interfaz-de-usuario)
4. [Gestión de Tareas](#gestión-de-tareas)
5. [Operaciones Comunes](#operaciones-comunes)
6. [Preguntas Frecuentes](#preguntas-frecuentes)
7. [Contacto y Soporte](#contacto-y-soporte)

---

## Introducción

### ¿Qué es esta aplicación?

Esta es una **aplicación web de gestión de tareas** que permite:
- ✅ Crear nuevas tareas
- ✅ Ver listado de tareas
- ✅ Marcar tareas como completadas
- ✅ Editar tareas existentes
- ✅ Eliminar tareas

La aplicación funciona en tu navegador y guarda todos los datos en una base de datos segura.

### Requisitos
- Navegador web moderno (Chrome, Firefox, Safari, Edge)
- Conexión a Internet (si usa servidor remoto)
- Acceso a la URL donde está instalada la app

---

## Primeros Pasos

### 1. Acceder a la Aplicación

Abre tu navegador web e ingresa a:

**Local**:
```
http://localhost:8080
```

**Servidor Remoto**:
```
http://[dirección-del-servidor]:8080
```

### 2. Pantalla de Inicio

Verás una página con:
- **Título**: "GraphQL CRUD con Go"
- **Barra de Entrada**: Donde escribes nuevas tareas
- **Listado de Tareas**: Muestra todas tus tareas guardadas

---

## Interfaz de Usuario

### Componentes Principales

```
┌─────────────────────────────────────────────────────┐
│  GraphQL CRUD con Go                                │
├─────────────────────────────────────────────────────┤
│                                                      │
│  ┌──────────────────────────────────────┐            │
│  │ Nombre de la tarea                   │ + Agregar │
│  └──────────────────────────────────────┘            │
│  ┌──────────────────────────────────────────────────┐│
│  │ Descripción de la tarea                          ││
│  └──────────────────────────────────────────────────┘│
│                                                      │
│  Tareas:                                             │
│  ─────────────────────────────────────────────────  │
│  □ Tarea 1                  [Editar] [Eliminar]     │
│  ☑ Tarea 2 (completada)     [Editar] [Eliminar]     │
│  □ Tarea 3                  [Editar] [Eliminar]     │
│                                                      │
└─────────────────────────────────────────────────────┘
```

### Elementos de la Interfaz

| Elemento | Descripción |
|----------|-------------|
| **Campo "Nombre de la tarea"** | Escribe aquí el título de la nueva tarea |
| **Campo "Descripción"** | (Opcional) Escribe detalles sobre la tarea |
| **Botón "+ Agregar"** | Crea la nueva tarea |
| **Casilla de verificación (□/☑)** | Marca para completar una tarea |
| **Botón "Editar"** | Modifica la tarea |
| **Botón "Eliminar"** | Elimina la tarea |
| **Listado de Tareas** | Muestra todas tus tareas |

---

## Gestión de Tareas

### Crear una Nueva Tarea

**Paso 1**: Escribe el nombre de la tarea
- Haz clic en el campo "Nombre de la tarea"
- Escribe el título (ej: "Comprar leche")

**Paso 2**: (Opcional) Agrega descripción
- Haz clic en el campo "Descripción"
- Escribe detalles (ej: "2 litros, leche descremada")

**Paso 3**: Crea la tarea
- Haz clic en el botón "+ Agregar"
- La tarea aparecerá en el listado

### Ver Listado de Tareas

El listado muestra todas tus tareas con:
- **Casilla**: Para marcar como completada
- **Nombre**: Título de la tarea
- **Descripción**: Detalles (si existen)
- **Botones de acción**: Editar o Eliminar

### Marcar Tarea como Completada

**Opción 1: Desde el listado**
- Haz clic en la casilla (□) junto a la tarea
- La tarea se marcará con ☑
- El listado se actualiza automáticamente

**Opción 2: Desde la edición**
- Haz clic en "Editar" 
- Marca la casilla "Completada"
- Guarda cambios

### Editar una Tarea

**Paso 1**: Haz clic en "Editar"
- Se abrirá un formulario de edición

**Paso 2**: Modifica los campos
- Nombre
- Descripción
- Estado de completada (☑)

**Paso 3**: Guarda cambios
- Haz clic en "Guardar"
- El listado se actualiza

### Eliminar una Tarea

**Opción 1: Botón Eliminar**
1. Haz clic en "Eliminar" junto a la tarea
2. Se pedirá confirmación
3. Haz clic en "Sí, eliminar"
4. La tarea desaparece del listado

**Importante**: Esta acción no se puede deshacer. Ten cuidado al eliminar.

---

## Operaciones Comunes

### Buscar una Tarea

Aunque no hay buscador, puedes:
1. Scroll por el listado
2. Usar Ctrl+F (o Cmd+F en Mac) para buscar en la página

### Ver Detalles de una Tarea

1. En el listado, verás el nombre y descripción
2. Si quieres más detalle, haz clic en "Editar"

### Cambiar Estado de Completada

**Método 1 (Rápido)**:
- Haz clic en la casilla □ / ☑

**Método 2 (Desde edición)**:
- Clic en "Editar"
- Marca/desmarca "Completada"
- Guarda

### Filtrar Tareas Completadas

Por el momento, no hay filtro automático. Puedes:
- Mirar visualmente (☑ = completada)
- Usar Ctrl+F para buscar y filtrar manualmente

### Borrar Todas las Tareas

No hay botón para borrar todo. Debes eliminar una por una con el botón "Eliminar".

---

## Preguntas Frecuentes

### P: ¿Se guardan mis tareas?
**R**: Sí, todas las tareas se guardan automáticamente en la base de datos. Incluso si cierras el navegador, tus tareas seguirán ahí cuando regreses.

### P: ¿Cuántas tareas puedo crear?
**R**: Prácticamente ilimitadas. La capacidad depende del espacio disponible en la base de datos (típicamente miles o millones de tareas).

### P: ¿Puedo compartir mis tareas con otros?
**R**: No en esta versión. Esta es una aplicación personal. Si es necesario, contacta al administrador.

### P: ¿Qué sucede si elimino una tarea por error?
**R**: Las tareas eliminadas no se pueden recuperar. Asegúrate de que realmente quieras eliminarla antes de confirmar.

### P: ¿Mi información está segura?
**R**: Tus datos se guardan en una base de datos segura en el servidor. La aplicación no compartirá tus datos con terceros.

### P: ¿Puedo editar una tarea completada?
**R**: Sí, puedes editar cualquier tarea, incluso si ya está marcada como completada. Haz clic en "Editar" y modifica lo que necesites.

### P: ¿Hay límite de caracteres para el nombre?
**R**: El nombre puede tener hasta 255 caracteres. La descripción puede ser más larga.

### P: ¿Funciona en móvil?
**R**: Sí, la aplicación es responsiva y funciona en teléfonos y tablets. Accede desde cualquier navegador móvil.

### P: ¿Qué navegadores soporta?
**R**: Chrome, Firefox, Safari, Edge y cualquier navegador moderno con soporte para JavaScript.

### P: ¿Por qué tarda en cargar?
**R**: La primera carga puede tardar algunos segundos. Las cargas posteriores son más rápidas. Si sigue siendo lenta, contáctanos.

### P: ¿Puedo descargar mis tareas?
**R**: No en esta versión. Puedes copiar y pegar manualmente o contactar al administrador.

---

## Operaciones Avanzadas

### Usar desde la Línea de Comandos (Técnico)

Si tienes acceso a terminal, puedes usar la API GraphQL directamente:

#### Crear tarea
```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createTask(input: {name: \"Mi Tarea\"}) { id name } }"
  }'
```

#### Listar tareas
```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ tasks { id name completed } }"}'
```

---

## Solución de Problemas

### La página no carga
**Posibles causas**:
- El servidor no está corriendo
- Ingresaste una URL incorrecta
- No hay conexión de Internet

**Solución**:
1. Verifica la URL correcta
2. Pide al administrador que reinicie la aplicación
3. Intenta en otro navegador

### Las tareas no se guardan
**Posibles causas**:
- Base de datos no está disponible
- Error de conexión
- Falta de permisos

**Solución**:
1. Recarga la página (Ctrl+R)
2. Cierra el navegador y vuelve a abrir
3. Contacta al administrador

### Botones no responden
**Posibles causas**:
- JavaScript deshabilitado en el navegador
- Página no cargó completamente
- Servidor sobrecargado

**Solución**:
1. Habilita JavaScript en tu navegador
2. Espera a que cargue completamente
3. Recarga la página
4. Intenta en otro navegador

### Se ve raro en mi dispositivo
**Solución**:
1. Intenta cambiar la orientación del dispositivo
2. Zoom de la página (Ctrl +/-)
3. Usa otro navegador o dispositivo

---

## Contacto y Soporte

### ¿Necesitas ayuda?

**Para reportar problemas**:
- 📧 Email: contacto@example.com
- 🐙 GitHub: https://github.com/rianeiromiron/graphqlgo
- 💬 Issues: https://github.com/rianeiromiron/graphqlgo/issues

### Información Técnica para el Soporte

Cuando reportes un problema, incluye:
- Sistema operativo (Windows, Mac, Linux)
- Navegador y versión
- Pasos para reproducir el problema
- Mensajes de error (si los hay)
- Screenshot (si es posible)

---

## Consejos y Buenas Prácticas

### ✅ Haz esto
- Escribe títulos claros y descriptivos
- Usa la descripción para detalles importantes
- Marca tareas completadas regularmente
- Revisa tu lista regularmente

### ❌ No hagas esto
- No esperes a eliminar accidentalmente (no hay deshacer)
- No compartas tu URL de acceso si es privada
- No intentes editar directamente el código (si no eres técnico)
- No ignores los mensajes de error

---

## Atajos de Teclado

| Atajo | Acción |
|-------|--------|
| **Ctrl + R** (o Cmd + R) | Recargar página |
| **Ctrl + F** (o Cmd + F) | Buscar en página |
| **Tab** | Navegar entre campos |
| **Enter** | Enviar formulario |
| **Esc** | Cancelar operación |

---

## Glosario

| Término | Significado |
|---------|-------------|
| **Tarea** | Una item o elemento que necesitas recordar o completar |
| **Completada** | Una tarea que ya fue hecha (marcada con ☑) |
| **Descripción** | Detalles adicionales sobre una tarea |
| **API** | Sistema para comunicarse con la aplicación |
| **GraphQL** | Lenguaje para consultar datos (detalles técnicos) |
| **Base de Datos** | Lugar donde se guardan tus tareas |

---

**Manual de Usuario - Versión 1.0**
**Última actualización**: 2026-08-15
**Aplicación**: GraphQL Go CRUD
