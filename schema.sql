CREATE TABLE user (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  bio TEXT
)

CREATE TABLE flow (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT
)

CREATE TABLE sub_flow (
  id INTEGER PRIMARY KEY,
  flow_id INTEGER NOT NULL,
  name TEXT NOT NULL,
  description TEXT
  FOREIGN KEY (flow_id) REFERENCES flow(id) ON DELETE CASCADE
)

CREATE TABLE node (
  id uuid PRIMARY KEY,
  sub_flow_id INTEGER NOT NULL,
  type TEXT NOT NULL,
  parent TEXT,
  position TEXT,
  styles TEXT,
  width INTEGER,
  height INTEGER,
  hidden INTEGER,
  description TEXT,
  FOREIGN KEY (sub_flow_id) REFERENCES sub_flow(id) ON DELETE CASCADE
)

CREATE TABLE edge (
  id uuid PRIMARY KEY,
  sub_flow_id INTEGER NOT NULL,
  source INTEGER NOT NULL,
  target INTEGER NOT NULL,
  type TEXT NOT NULL,
  label TEXT,
  hidden INTEGER,
  marker_end TEXT,
  points TEXT,
  FOREIGN KEY (sub_flow_id) REFERENCES edge(id) ON DELETE CASCADE
)
