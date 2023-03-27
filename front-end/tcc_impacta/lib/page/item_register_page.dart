import 'package:flutter/material.dart';
class ItemRegister extends StatefulWidget {
  final int? id;
  final String? name;
  final DateTime? entryDate;
  final int? quantity;
  final DateTime? expirationDate;

  const ItemRegister({
    Key? key,
    required this.id,
    this.name,
    required this.entryDate,
    required this.quantity,
    required this.expirationDate,
  }) : super(key: key);

  @override
  _ItemRegisterState createState() => _ItemRegisterState();
}

class _ItemRegisterState extends State<ItemRegister> {
  final _formKey = GlobalKey<FormState>();

  String? _name;
  DateTime? _entryDate;
  int? _quantity;
  DateTime? _expirationDate;

  @override
  void initState() {
    super.initState();
    _name = widget.name;
    _entryDate = widget.entryDate;
    _quantity = widget.quantity;
    _expirationDate = widget.expirationDate;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: Card(
              color: Colors.grey[500],
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: Form(
                  key: _formKey,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      TextFormField(
                        initialValue: _name,
                        decoration: const InputDecoration(
                          labelText: 'Nome',
                        ),
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Please enter a name';
                          }
                          return null;
                        },
                        onSaved: (value) {
                          _name = value;
                        },
                      ),
                      TextFormField(
                        initialValue: _entryDate.toString(),
                        decoration: const InputDecoration(
                          labelText: 'Entry Date',
                        ),
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Please enter an entry date';
                          }
                          return null;
                        },
                        onSaved: (value) {
                          _entryDate = DateTime.parse(value!);
                        },
                      ),
                      TextFormField(
                        initialValue: _quantity.toString(),
                        decoration: const InputDecoration(
                          labelText: 'Quantity',
                        ),
                        keyboardType: TextInputType.number,
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Please enter a quantity';
                          }
                          if (int.tryParse(value) == null) {
                            return 'Please enter a valid number';
                          }
                          return null;
                        },
                        onSaved: (value) {
                          _quantity = int.parse(value!);
                        },
                      ),
                      TextFormField(
                        initialValue: _expirationDate.toString(),
                        decoration: const InputDecoration(
                          labelText: 'Expiration Date',
                        ),
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Please enter an expiration date';
                          }
                          return null;
                        },
                        onSaved: (value) {
                          _expirationDate = DateTime.parse(value!);
                        },
                      ),
                      const SizedBox(height: 16),
                      Center(
                        child: ElevatedButton(
                          onPressed: () {
                            if (_formKey.currentState!.validate()) {
                              _formKey.currentState!.save();
                            }
                            setState(() {});
                          },
                          child: const Text('Save'),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
