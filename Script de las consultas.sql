create database bases1_p1;

use bases1_p1;

INSERT INTO orden
select distinct id_orden,fecha_orden,id_cliente
from temp_orden;

INSERT INTO detorden(cantidad, id_producto, id_orden, id_vendedor)
SELECT cantidad, id_producto, id_orden, id_vendedor from temp_orden;


-- CONSULTS


-- Consult 1
WITH sup as ( select cliente.id_cliente, concat(cliente.nombre,' ',cliente.apellido) as NombreCompleto, p.nombre as Pais, p2.precio*d.cantidad as Total from cliente
                       inner join orden o on cliente.id_cliente = o.id_cliente
                       inner join detorden d on o.id_orden = d.id_orden
                       inner join pais p on cliente.id_pais = p.id_pais
                       inner join producto p2 on d.id_producto = p2.id_producto)
select sup.id_cliente as ID_Cliente, sup.NombreCompleto, sup.Pais ,SUM(sup.Total) as Monto_Total
from sup group by sup.NombreCompleto order by Monto_Total DESC limit 1;

-- Consult 2

with sup1 as (SELECT p.id_producto, p.nombre as Nombre_Producto, c.Nombre as Categoria, SUM(d.cantidad) as Cantidad, SUM(p.precio*d.cantidad) as precio FROM producto as p
                       inner join categoria c on p.id_categoria = c.id_categoria
                       inner join detorden d on p.id_producto = d.id_producto
                       group by Nombre_Producto )
(SELECT * from sup1 order by sup1.Cantidad, sup1.precio limit 1)
UNION ALL
(SELECT * from sup1 order by sup1.Cantidad desc , sup1.precio limit 1);

-- Consult 3

WITH sup as (SELECT v.id_vendedor, v.nombre, d.cantidad, p.precio as precio_uni ,p.precio*d.cantidad as precio from vendedor as v
                      inner join detorden d on v.id_vendedor = d.id_vendedor
                      inner join producto p on d.id_producto = p.id_producto)
SELECT sup.id_vendedor, sup.nombre, SUM(sup.precio) as Vendido
from sup group by sup.nombre order by Vendido  DESC limit 1;

-- Consult 4


with sup1 as (SELECT p2.nombre, SUM(p.precio*d.cantidad) as precio from vendedor as v
                      inner join detorden d on v.id_vendedor = d.id_vendedor
                      inner join producto p on d.id_producto = p.id_producto
                      inner join pais p2 on v.id_pais = p2.id_pais
                      group by p2.nombre)
(SELECT * from sup1 order by sup1.precio desc limit 1)
UNION ALL
(SELECT * from sup1 order by sup1.precio limit 1);


-- Consult 5

WITH sup as (SELECT pais.id_pais, pais.nombre, p.precio*d.cantidad as precio from pais
                      inner join cliente c on pais.id_pais = c.id_pais
                      inner join orden o on c.id_cliente = o.id_cliente
                      inner join detorden d on o.id_orden = d.id_orden
                      inner join producto p on d.id_producto = p.id_producto)
SELECT sup.id_pais, sup.nombre as Pais, SUM(sup.precio) as Monto
FROM sup group by Pais order by Monto
limit 5;

-- Consult 6

with sup1 as(SELECT c.Nombre, SUM(d.cantidad) as Cantidad from categoria as c
                     inner join producto p on c.id_categoria = p.id_categoria
                     inner join detorden d on p.id_producto = d.id_producto
                     group by c.nombre)
(SELECT * from sup1 order by sup1.Cantidad desc limit 1)
UNION ALL
(SELECT * from sup1 order by sup1.Cantidad limit 1);

-- Consult 7
with sup as(SELECT p2.nombre as pais, c.Nombre as cate, SUM(d.cantidad) as total from categoria as c
inner join producto p on c.id_categoria = p.id_categoria
inner join detorden d on p.id_producto = d.id_producto
inner join orden o on d.id_orden = o.id_orden
inner join cliente c2 on o.id_cliente = c2.id_cliente
inner join pais p2 on c2.id_pais = p2.id_pais
group by  pais, cate order by total desc)
SELECT sup.pais as Pais, sup.cate as Categoria, sup.total as Cantidad from sup
inner join
    (SELECT pais,cate, MAX(total) as Total from sup group by pais) sup2
on sup.pais=sup2.pais
and sup.cate = sup2.cate;

-- Consult 8

SELECT MONTH(o.fecha) as Mes, SUM(p.precio*d.cantidad) as Total from pais
    inner join vendedor v on pais.id_pais = v.id_pais
    inner join detorden d on v.id_vendedor = d.id_vendedor
    inner join orden o on d.id_orden = o.id_orden
    inner join producto p on d.id_producto = p.id_producto
    where pais.nombre = 'Inglaterra'
    group by month(o.fecha);

-- Consult 9


with sup1 as (SELECT MONTH(o.fecha) as Mes, SUM(p.precio*d.cantidad) as Total from pais
    inner join vendedor v on pais.id_pais = v.id_pais
    inner join detorden d on v.id_vendedor = d.id_vendedor
    inner join orden o on d.id_orden = o.id_orden
    inner join producto p on d.id_producto = p.id_producto
    group by month(o.fecha) )
(SELECT * from sup1 order by sup1.Total desc limit 1)
UNION ALL
(SELECT * from sup1 order by sup1.Total limit 1);

-- Consult 10
SELECT p.id_producto as 'Id_Producto', p.nombre as 'Nombre Producto' ,SUM(p.precio*d.cantidad) as Total from pais
    inner join vendedor v on pais.id_pais = v.id_pais
    inner join detorden d on v.id_vendedor = d.id_vendedor
    inner join orden o on d.id_orden = o.id_orden
    inner join producto p on d.id_producto = p.id_producto
    inner join categoria c on p.id_categoria = c.id_categoria
    WHERE c.Nombre = 'Deportes'
    group by `Nombre Producto` order by Total;