class User {
  final String? id;  // Make id nullable
  final String username;
  String? password;

  User({
    this.id,  // Remove required
    required this.username,
    this.password,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id']?.toString(),  // Handle potential null and convert to String
      username: json['username'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'username': username,
      'password': password,
    };
  }
}