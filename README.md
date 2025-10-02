# Documentation: Manipulation de MongoDB avec Go

## **1. Configuration de MongoDB**
- **Docker Compose** : MongoDB est configuré dans le fichier `docker-compose.yml`. Pour démarrer MongoDB :
  ```bash
  docker-compose up -d
  ```
- MongoDB est accessible à l'adresse : `mongodb://root:example@localhost:27017`.

---

## **2. Fichier `conexion.go`**
Ce fichier contient la logique principale pour interagir avec MongoDB.

### **Fonctions disponibles :**

1. **`ConnectMongoDB()`**
   - Établit une connexion avec MongoDB.
   - Vérifie la connexion avec un ping.
   - Exemple d'utilisation :
     ```go
     ConnectMongoDB()
     ```

2. **`AddUser(username string, password string)`**
   - Ajoute un utilisateur dans la collection `users` de la base de données `testdb`.
   - Exemple d'utilisation :
     ```go
     AddUser("testuser", "testpassword")
     ```

3. **`Login(username string, password string) bool`**
   - Valide les informations d'identification d'un utilisateur.
   - Retourne `true` si les informations sont correctes, sinon `false`.
   - Exemple d'utilisation :
     ```go
     if Login("testuser", "testpassword") {
         fmt.Println("Login successful")
     } else {
         fmt.Println("Login failed")
     }
     ```

---

## **3. Fichier `main.go`**
Ce fichier est utilisé pour tester la logique MongoDB.

### **Exemple de flux :**
1. Connexion à MongoDB.
2. Ajout d'un utilisateur.
3. Validation de l'utilisateur.

Exemple de code dans `main.go` :
```go
package main

func main() {
    // Connexion à MongoDB
    ConnectMongoDB()

    // Ajout d'un utilisateur
    AddUser("testuser", "testpassword")

    // Test de connexion
    if Login("testuser", "testpassword") {
        fmt.Println("Login successful")
    } else {
        fmt.Println("Login failed")
    }
}
```

---

## **4. Manipulation des données avec MongoDB Shell**
Vous pouvez également manipuler les données directement dans MongoDB via le shell `mongosh`.

### **Commandes utiles :**
1. **Connexion à MongoDB :**
   ```bash
   docker run -it --rm --network host mongo mongosh mongodb://root:example@localhost:27017
   ```

2. **Changer de base de données :**
   ```javascript
   use testdb
   ```

3. **Afficher les utilisateurs :**
   ```javascript
   db.users.find()
   ```

4. **Ajouter un utilisateur :**
   ```javascript
   db.users.insertOne({ username: "newuser", password: "newpassword" })
   ```

5. **Supprimer un utilisateur :**
   ```javascript
   db.users.deleteOne({ username: "testuser" })
   ```

---

## **5. Dépannage**
- **MongoDB ne démarre pas :**
  - Vérifiez que Docker est en cours d'exécution.
  - Utilisez `docker-compose logs mongodb` pour voir les erreurs.

- **Erreur de connexion dans Go :**
  - Assurez-vous que MongoDB est en cours d'exécution (`docker ps`).
  - Vérifiez l'URI de connexion dans `ConnectMongoDB`.

- **Login échoue :**
  - Vérifiez que l'utilisateur existe dans la collection `users`.
  - Utilisez `db.users.find()` pour confirmer.

---

## **6. Étendre la logique**
Vous pouvez ajouter d'autres fonctions dans `conexion.go` pour manipuler les données :

### **Mise à jour d'un utilisateur :**
```go
func UpdateUser(username, newPassword string) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.UpdateOne(ctx, map[string]interface{}{
        "username": username,
    }, map[string]interface{}{
        "$set": map[string]interface{}{
            "password": newPassword,
        },
    })
    if err != nil {
        log.Fatalf("Failed to update user: %v", err)
    }

    fmt.Println("User updated successfully")
}
```

---

## **7. Commandes Résumées**

### **Installation des dépendances Go :**
```bash
go get go.mongodb.org/mongo-driver/mongo/options
```

### **Démarrer MongoDB avec Docker :**
```bash
docker-compose up -d
```

### **Exécuter l'application Go :**
```bash
go run main.go conexion.go
```

---

Si vous avez besoin d'aide supplémentaire, n'hésitez pas à demander !