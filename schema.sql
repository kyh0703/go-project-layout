CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  bio TEXT
)

CREATE TABLE flows (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT
)

CREATE TABLE sub_flows (
  id INTEGER PRIMARY KEY,
  flow_id INTEGER NOT NULL,
  name TEXT NOT NULL,
  description TEXT
  FOREIGN KEY (flow_id) REFERENCES flows(id) ON DELETE CASCADE
)

CREATE TABLE nodes (
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
  FOREIGN KEY (sub_flow_id) REFERENCES sub_flows(id) ON DELETE CASCADE
)

CREATE TABLE edges (
  id uuid PRIMARY KEY,
  sub_flow_id INTEGER NOT NULL,
  source INTEGER NOT NULL,
  target INTEGER NOT NULL,
  type TEXT NOT NULL,
  label TEXT,
  hidden INTEGER,
  marker_end TEXT,
  points TEXT,
  FOREIGN KEY (sub_flow_id) REFERENCES edges(id) ON DELETE CASCADE
)
