INSERT INTO pilots (id, name) VALUES
(1, 'Manfred'),
(2, 'William'),
(3, 'Albert'),
(4, 'Warner'),
(5, 'Andrew');

INSERT INTO languages (id, language) VALUES
(1, 'German'),
(2, 'English');

INSERT INTO pilot_languages (pilot_id, language_id) VALUES
(1, 1),
(2, 2),
(3, 2),
(4, 1),
(4, 2);

INSERT INTO jets (id, pilot_id, age, name, color) VALUES
(1, 1, 10, 'F-1', 'White'),
(2, 2, 25, 'F-2', 'White'),
(3, 3, 25, 'F-3', 'White'),
(4, 4, 20, 'F-4', 'White'),
(5, 5, 15, 'F-5', 'White'),
(6, 1, 25, 'F-6', 'White'),
(7, 2, 10, 'F-7', 'White'),
(8, 3, 1, 'F-8', 'White'),
(9, 4, 30, 'F-9', 'White'),
(10, 5, 10, 'F-10', 'White'),
(11, 1, 20, 'F-11', 'White'),
(12, 2, 10, 'F-12', 'White'),
(13, 3, 30, 'F-13', 'White'),
(14, 4, 10, 'F-14', 'White'),
(15, 5, 40, 'F-15', 'White'),
(16, 1, 10, 'F-16', 'White'),
(17, 2, 10, 'F-17', 'White'),
(18, 3, 20, 'F-18', 'White'),
(19, 4, 30, 'F-19', 'White'),
(20, 5, 10, 'F-20', 'White'),
(21, 1, 25, 'F-21', 'White');
